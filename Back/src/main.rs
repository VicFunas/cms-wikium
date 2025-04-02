use tonic::{transport::Server, Request, Response, Status};

pub mod wiki {
    tonic::include_proto!("wiki");
}

use wiki::{
    wiki_service_server::{WikiService, WikiServiceServer},
    CreatePageRequest, ListPagesRequest, ListPagesResponse, ListVersionsResponse, PageRequest,
    PageResponse, PageVersionsRequest, UpdatePageRequest, VersionRequest,
};

#[derive(Default)]
pub struct WikiServiceImpl {}

#[tonic::async_trait]
impl WikiService for WikiServiceImpl {
    async fn get_page(
        &self,
        request: Request<PageRequest>,
    ) -> Result<Response<PageResponse>, Status> {
        // TODO: Implement with Fauna
        unimplemented!()
    }

    async fn create_page(
        &self,
        request: Request<CreatePageRequest>,
    ) -> Result<Response<PageResponse>, Status> {
        // TODO: Implement with Fauna
        unimplemented!()
    }

    async fn update_page(
        &self,
        request: Request<UpdatePageRequest>,
    ) -> Result<Response<PageResponse>, Status> {
        // TODO: Implement with Fauna
        unimplemented!()
    }

    async fn list_pages(
        &self,
        request: Request<ListPagesRequest>,
    ) -> Result<Response<ListPagesResponse>, Status> {
        // TODO: Implement with Fauna
        unimplemented!()
    }

    async fn get_page_version(
        &self,
        request: Request<VersionRequest>,
    ) -> Result<Response<PageResponse>, Status> {
        // TODO: Implement with Fauna
        unimplemented!()
    }

    async fn list_page_versions(
        &self,
        request: Request<PageVersionsRequest>,
    ) -> Result<Response<ListVersionsResponse>, Status> {
        // TODO: Implement with Fauna
        unimplemented!()
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse()?;
    let service = WikiServiceImpl::default();

    println!("Wiki gRPC server listening on {}", addr);

    Server::builder()
        .add_service(WikiServiceServer::new(service))
        .serve(addr)
        .await?;

    Ok(())
}
