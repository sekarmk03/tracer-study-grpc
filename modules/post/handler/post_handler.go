package handler

import (
	"context"
	"log"
	"net/http"
	"tracer-study-grpc/common/config"
	"tracer-study-grpc/common/errors"
	"tracer-study-grpc/modules/post/entity"
	"tracer-study-grpc/modules/post/service"
	"tracer-study-grpc/pb"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostHandler struct {
	pb.UnimplementedPostServiceServer
	config   config.Config
	postSvc  service.PostServiceUseCase
	imageSvc service.ImageServiceUseCase
}

func NewPostHandler(config config.Config, postService service.PostServiceUseCase, imageService service.ImageServiceUseCase) *PostHandler {
	return &PostHandler{
		config:   config,
		postSvc:  postService,
		imageSvc: imageService,
	}
}

func (ph *PostHandler) GetAllPosts(ctx context.Context, req *emptypb.Empty) (*pb.GetAllPostsResponse, error) {
	post, err := ph.postSvc.FindAll(ctx, req)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - GetAllPost] Error while get all post: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	var postArr []*pb.Post
	for _, p := range post {
		postProto := entity.ConvertEntityToProto(p)
		postArr = append(postArr, postProto)
	}

	return &pb.GetAllPostsResponse{
		Code:    uint32(http.StatusOK),
		Message: "get all post success",
		Data:    postArr,
	}, nil
}

func (ph *PostHandler) GetPostById(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.GetPostResponse, error) {
	post, err := ph.postSvc.FindById(ctx, req.GetId())
	if err != nil {
		if post == nil {
			log.Println("WARNING: [PostHandler - GetPostById] Resource post not found for id:", req.GetId())
			return nil, status.Errorf(status.Code(err), "post not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - GetPostById] Internal server error:", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	postProto := entity.ConvertEntityToProto(post)
	return &pb.GetPostResponse{
		Code:    uint32(http.StatusOK),
		Message: "get post success",
		Data:    postProto,
	}, nil
}

func (ph *PostHandler) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.GetPostResponse, error) {
	image, err := ph.imageSvc.UploadImage(ctx, req.GetImageFilename(), req.GetImageBuffer())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - CreatePost] Error while upload image: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	post, err := ph.postSvc.Create(ctx, req.GetTitle(), req.GetContent(), image, req.GetImageCaption(), req.GetType(), req.GetIsFeatured(), req.GetCreatedBy(), req.GetTags())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - CreatePost] Error while create post: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	postProto := entity.ConvertEntityToProto(post)
	return &pb.GetPostResponse{
		Code:    uint32(http.StatusOK),
		Message: "create post success",
		Data:    postProto,
	}, nil
}

func (ph *PostHandler) UpdatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.GetPostResponse, error) {
	post, err := ph.postSvc.FindById(ctx, req.GetId())
	if err != nil {
		if post == nil {
			log.Println("WARNING: [PostHandler - GetPostById] Resource post not found for id:", req.GetId())
			return nil, status.Errorf(status.Code(err), "post not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - GetPostById] Internal server error:", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	image, err := ph.imageSvc.UploadImage(ctx, req.GetImageFilename(), req.GetImageBuffer())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - UpdatePost] Error while upload image: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	err = ph.imageSvc.DeleteImage(ctx, post.MainImage)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - UpdatePost] Error while delete image: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	postDataUpdate := &entity.Post{
		Title:        req.GetTitle(),
		Content:      req.GetContent(),
		MainImage:    image,
		CaptionImage: req.GetImageCaption(),
		Type:         req.GetType(),
		IsFeatured:   req.GetIsFeatured(),
		CreatedBy:    req.GetCreatedBy(),
		UpdatedBy:    req.GetUpdatedBy(),
		Tags:         req.GetTags(),
	}

	post, err = ph.postSvc.Update(ctx, req.GetId(), postDataUpdate)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - UpdatePost] Error while update post: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	postProto := entity.ConvertEntityToProto(post)
	return &pb.GetPostResponse{
		Code:    uint32(http.StatusOK),
		Message: "update post success",
		Data:    postProto,
	}, nil
}

func (ph *PostHandler) DeletePost(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.DeletePostResponse, error) {
	post, err := ph.postSvc.FindById(ctx, req.GetId())
	if err != nil {
		if post == nil {
			log.Println("WARNING: [PostHandler - GetPostById] Resource post not found for id:", req.GetId())
			return nil, status.Errorf(status.Code(err), "post not found")
		}
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - GetPostById] Internal server error:", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	err = ph.imageSvc.DeleteImage(ctx, post.MainImage)
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - DeletePost] Error while delete image: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	err = ph.postSvc.Delete(ctx, req.GetId())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - DeletePost] Error while delete post: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	return &pb.DeletePostResponse{
		Code:    uint32(http.StatusOK),
		Message: "delete post success",
	}, nil
}

func (ph *PostHandler) AddVisitor(ctx context.Context, req *pb.GetPostByIdRequest) (*pb.GetPostResponse, error) {
	post, err := ph.postSvc.IncrementVisitor(ctx, req.GetId())
	if err != nil {
		parseError := errors.ParseError(err)
		log.Println("ERROR: [PostHandler - AddVisitor] Error while increment visitor: ", parseError.Message)
		return nil, status.Errorf(parseError.Code, parseError.Message)
	}

	postProto := entity.ConvertEntityToProto(post)
	return &pb.GetPostResponse{
		Code:    uint32(http.StatusOK),
		Message: "increment visitor success",
		Data:    postProto,
	}, nil
}
