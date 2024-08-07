// @generated
/// Generated client implementations.
pub mod advertisers_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct AdvertisersServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl AdvertisersServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> AdvertisersServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> AdvertisersServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            AdvertisersServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn list_advertisers(
            &mut self,
            request: impl tonic::IntoRequest<super::ListAdvertisersRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListAdvertisersResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.AdvertisersService/ListAdvertisers",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.AdvertisersService", "ListAdvertisers"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_advertiser(
            &mut self,
            request: impl tonic::IntoRequest<super::GetAdvertiserRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetAdvertiserResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.AdvertisersService/GetAdvertiser",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.AdvertisersService", "GetAdvertiser"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn create_advertiser(
            &mut self,
            request: impl tonic::IntoRequest<super::CreateAdvertiserRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateAdvertiserResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.AdvertisersService/CreateAdvertiser",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.AdvertisersService", "CreateAdvertiser"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn update_advertiser(
            &mut self,
            request: impl tonic::IntoRequest<super::UpdateAdvertiserRequest>,
        ) -> std::result::Result<
            tonic::Response<super::UpdateAdvertiserResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.AdvertisersService/UpdateAdvertiser",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.AdvertisersService", "UpdateAdvertiser"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn delete_advertiser(
            &mut self,
            request: impl tonic::IntoRequest<super::DeleteAdvertiserRequest>,
        ) -> std::result::Result<
            tonic::Response<super::DeleteAdvertiserResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.AdvertisersService/DeleteAdvertiser",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.AdvertisersService", "DeleteAdvertiser"),
                );
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod advertisers_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with AdvertisersServiceServer.
    #[async_trait]
    pub trait AdvertisersService: Send + Sync + 'static {
        async fn list_advertisers(
            &self,
            request: tonic::Request<super::ListAdvertisersRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListAdvertisersResponse>,
            tonic::Status,
        >;
        async fn get_advertiser(
            &self,
            request: tonic::Request<super::GetAdvertiserRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetAdvertiserResponse>,
            tonic::Status,
        >;
        async fn create_advertiser(
            &self,
            request: tonic::Request<super::CreateAdvertiserRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateAdvertiserResponse>,
            tonic::Status,
        >;
        async fn update_advertiser(
            &self,
            request: tonic::Request<super::UpdateAdvertiserRequest>,
        ) -> std::result::Result<
            tonic::Response<super::UpdateAdvertiserResponse>,
            tonic::Status,
        >;
        async fn delete_advertiser(
            &self,
            request: tonic::Request<super::DeleteAdvertiserRequest>,
        ) -> std::result::Result<
            tonic::Response<super::DeleteAdvertiserResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct AdvertisersServiceServer<T: AdvertisersService> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T: AdvertisersService> AdvertisersServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for AdvertisersServiceServer<T>
    where
        T: AdvertisersService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/tap.ads.v1.AdvertisersService/ListAdvertisers" => {
                    #[allow(non_camel_case_types)]
                    struct ListAdvertisersSvc<T: AdvertisersService>(pub Arc<T>);
                    impl<
                        T: AdvertisersService,
                    > tonic::server::UnaryService<super::ListAdvertisersRequest>
                    for ListAdvertisersSvc<T> {
                        type Response = super::ListAdvertisersResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListAdvertisersRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as AdvertisersService>::list_advertisers(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ListAdvertisersSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.AdvertisersService/GetAdvertiser" => {
                    #[allow(non_camel_case_types)]
                    struct GetAdvertiserSvc<T: AdvertisersService>(pub Arc<T>);
                    impl<
                        T: AdvertisersService,
                    > tonic::server::UnaryService<super::GetAdvertiserRequest>
                    for GetAdvertiserSvc<T> {
                        type Response = super::GetAdvertiserResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GetAdvertiserRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as AdvertisersService>::get_advertiser(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = GetAdvertiserSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.AdvertisersService/CreateAdvertiser" => {
                    #[allow(non_camel_case_types)]
                    struct CreateAdvertiserSvc<T: AdvertisersService>(pub Arc<T>);
                    impl<
                        T: AdvertisersService,
                    > tonic::server::UnaryService<super::CreateAdvertiserRequest>
                    for CreateAdvertiserSvc<T> {
                        type Response = super::CreateAdvertiserResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CreateAdvertiserRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as AdvertisersService>::create_advertiser(
                                        &inner,
                                        request,
                                    )
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = CreateAdvertiserSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.AdvertisersService/UpdateAdvertiser" => {
                    #[allow(non_camel_case_types)]
                    struct UpdateAdvertiserSvc<T: AdvertisersService>(pub Arc<T>);
                    impl<
                        T: AdvertisersService,
                    > tonic::server::UnaryService<super::UpdateAdvertiserRequest>
                    for UpdateAdvertiserSvc<T> {
                        type Response = super::UpdateAdvertiserResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::UpdateAdvertiserRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as AdvertisersService>::update_advertiser(
                                        &inner,
                                        request,
                                    )
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = UpdateAdvertiserSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.AdvertisersService/DeleteAdvertiser" => {
                    #[allow(non_camel_case_types)]
                    struct DeleteAdvertiserSvc<T: AdvertisersService>(pub Arc<T>);
                    impl<
                        T: AdvertisersService,
                    > tonic::server::UnaryService<super::DeleteAdvertiserRequest>
                    for DeleteAdvertiserSvc<T> {
                        type Response = super::DeleteAdvertiserResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::DeleteAdvertiserRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as AdvertisersService>::delete_advertiser(
                                        &inner,
                                        request,
                                    )
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = DeleteAdvertiserSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", tonic::Code::Unimplemented as i32)
                                .header(
                                    http::header::CONTENT_TYPE,
                                    tonic::metadata::GRPC_CONTENT_TYPE,
                                )
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: AdvertisersService> Clone for AdvertisersServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: AdvertisersService> tonic::server::NamedService
    for AdvertisersServiceServer<T> {
        const NAME: &'static str = "tap.ads.v1.AdvertisersService";
    }
}
/// Generated client implementations.
pub mod requirements_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct RequirementsServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl RequirementsServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> RequirementsServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> RequirementsServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            RequirementsServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn list_requirements(
            &mut self,
            request: impl tonic::IntoRequest<super::ListRequirementsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListRequirementsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.RequirementsService/ListRequirements",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.RequirementsService", "ListRequirements"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn create_requirement(
            &mut self,
            request: impl tonic::IntoRequest<super::CreateRequirementRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateRequirementResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.RequirementsService/CreateRequirement",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new(
                        "tap.ads.v1.RequirementsService",
                        "CreateRequirement",
                    ),
                );
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod requirements_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with RequirementsServiceServer.
    #[async_trait]
    pub trait RequirementsService: Send + Sync + 'static {
        async fn list_requirements(
            &self,
            request: tonic::Request<super::ListRequirementsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListRequirementsResponse>,
            tonic::Status,
        >;
        async fn create_requirement(
            &self,
            request: tonic::Request<super::CreateRequirementRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateRequirementResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct RequirementsServiceServer<T: RequirementsService> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T: RequirementsService> RequirementsServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for RequirementsServiceServer<T>
    where
        T: RequirementsService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/tap.ads.v1.RequirementsService/ListRequirements" => {
                    #[allow(non_camel_case_types)]
                    struct ListRequirementsSvc<T: RequirementsService>(pub Arc<T>);
                    impl<
                        T: RequirementsService,
                    > tonic::server::UnaryService<super::ListRequirementsRequest>
                    for ListRequirementsSvc<T> {
                        type Response = super::ListRequirementsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListRequirementsRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as RequirementsService>::list_requirements(
                                        &inner,
                                        request,
                                    )
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ListRequirementsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.RequirementsService/CreateRequirement" => {
                    #[allow(non_camel_case_types)]
                    struct CreateRequirementSvc<T: RequirementsService>(pub Arc<T>);
                    impl<
                        T: RequirementsService,
                    > tonic::server::UnaryService<super::CreateRequirementRequest>
                    for CreateRequirementSvc<T> {
                        type Response = super::CreateRequirementResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CreateRequirementRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as RequirementsService>::create_requirement(
                                        &inner,
                                        request,
                                    )
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = CreateRequirementSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", tonic::Code::Unimplemented as i32)
                                .header(
                                    http::header::CONTENT_TYPE,
                                    tonic::metadata::GRPC_CONTENT_TYPE,
                                )
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: RequirementsService> Clone for RequirementsServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: RequirementsService> tonic::server::NamedService
    for RequirementsServiceServer<T> {
        const NAME: &'static str = "tap.ads.v1.RequirementsService";
    }
}
/// Generated client implementations.
pub mod passes_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct PassesServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl PassesServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> PassesServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> PassesServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            PassesServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn list_passes(
            &mut self,
            request: impl tonic::IntoRequest<super::ListPassesRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListPassesResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.PassesService/ListPasses",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.PassesService", "ListPasses"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_pass(
            &mut self,
            request: impl tonic::IntoRequest<super::GetPassRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetPassResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.PassesService/GetPass",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.PassesService", "GetPass"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn create_pass(
            &mut self,
            request: impl tonic::IntoRequest<super::CreatePassRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreatePassResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.PassesService/CreatePass",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.PassesService", "CreatePass"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn update_pass(
            &mut self,
            request: impl tonic::IntoRequest<super::UpdatePassRequest>,
        ) -> std::result::Result<
            tonic::Response<super::UpdatePassResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.PassesService/UpdatePass",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.PassesService", "UpdatePass"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn delete_pass(
            &mut self,
            request: impl tonic::IntoRequest<super::DeletePassRequest>,
        ) -> std::result::Result<
            tonic::Response<super::DeletePassResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.PassesService/DeletePass",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.PassesService", "DeletePass"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn list_pass_requirements(
            &mut self,
            request: impl tonic::IntoRequest<super::ListPassRequirementsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListPassRequirementsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.PassesService/ListPassRequirements",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.PassesService", "ListPassRequirements"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn create_pass_requirement(
            &mut self,
            request: impl tonic::IntoRequest<super::CreatePassRequirementRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreatePassRequirementResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.PassesService/CreatePassRequirement",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.PassesService", "CreatePassRequirement"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn delete_pass_requirement(
            &mut self,
            request: impl tonic::IntoRequest<super::DeletePassRequirementRequest>,
        ) -> std::result::Result<
            tonic::Response<super::DeletePassRequirementResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.PassesService/DeletePassRequirement",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.PassesService", "DeletePassRequirement"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn can_subscribe_to_pass(
            &mut self,
            request: impl tonic::IntoRequest<super::CanSubscribeToPassRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CanSubscribeToPassResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.PassesService/CanSubscribeToPass",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.PassesService", "CanSubscribeToPass"),
                );
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod passes_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with PassesServiceServer.
    #[async_trait]
    pub trait PassesService: Send + Sync + 'static {
        async fn list_passes(
            &self,
            request: tonic::Request<super::ListPassesRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListPassesResponse>,
            tonic::Status,
        >;
        async fn get_pass(
            &self,
            request: tonic::Request<super::GetPassRequest>,
        ) -> std::result::Result<tonic::Response<super::GetPassResponse>, tonic::Status>;
        async fn create_pass(
            &self,
            request: tonic::Request<super::CreatePassRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreatePassResponse>,
            tonic::Status,
        >;
        async fn update_pass(
            &self,
            request: tonic::Request<super::UpdatePassRequest>,
        ) -> std::result::Result<
            tonic::Response<super::UpdatePassResponse>,
            tonic::Status,
        >;
        async fn delete_pass(
            &self,
            request: tonic::Request<super::DeletePassRequest>,
        ) -> std::result::Result<
            tonic::Response<super::DeletePassResponse>,
            tonic::Status,
        >;
        async fn list_pass_requirements(
            &self,
            request: tonic::Request<super::ListPassRequirementsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListPassRequirementsResponse>,
            tonic::Status,
        >;
        async fn create_pass_requirement(
            &self,
            request: tonic::Request<super::CreatePassRequirementRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreatePassRequirementResponse>,
            tonic::Status,
        >;
        async fn delete_pass_requirement(
            &self,
            request: tonic::Request<super::DeletePassRequirementRequest>,
        ) -> std::result::Result<
            tonic::Response<super::DeletePassRequirementResponse>,
            tonic::Status,
        >;
        async fn can_subscribe_to_pass(
            &self,
            request: tonic::Request<super::CanSubscribeToPassRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CanSubscribeToPassResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct PassesServiceServer<T: PassesService> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T: PassesService> PassesServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for PassesServiceServer<T>
    where
        T: PassesService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/tap.ads.v1.PassesService/ListPasses" => {
                    #[allow(non_camel_case_types)]
                    struct ListPassesSvc<T: PassesService>(pub Arc<T>);
                    impl<
                        T: PassesService,
                    > tonic::server::UnaryService<super::ListPassesRequest>
                    for ListPassesSvc<T> {
                        type Response = super::ListPassesResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListPassesRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as PassesService>::list_passes(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ListPassesSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.PassesService/GetPass" => {
                    #[allow(non_camel_case_types)]
                    struct GetPassSvc<T: PassesService>(pub Arc<T>);
                    impl<
                        T: PassesService,
                    > tonic::server::UnaryService<super::GetPassRequest>
                    for GetPassSvc<T> {
                        type Response = super::GetPassResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GetPassRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as PassesService>::get_pass(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = GetPassSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.PassesService/CreatePass" => {
                    #[allow(non_camel_case_types)]
                    struct CreatePassSvc<T: PassesService>(pub Arc<T>);
                    impl<
                        T: PassesService,
                    > tonic::server::UnaryService<super::CreatePassRequest>
                    for CreatePassSvc<T> {
                        type Response = super::CreatePassResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CreatePassRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as PassesService>::create_pass(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = CreatePassSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.PassesService/UpdatePass" => {
                    #[allow(non_camel_case_types)]
                    struct UpdatePassSvc<T: PassesService>(pub Arc<T>);
                    impl<
                        T: PassesService,
                    > tonic::server::UnaryService<super::UpdatePassRequest>
                    for UpdatePassSvc<T> {
                        type Response = super::UpdatePassResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::UpdatePassRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as PassesService>::update_pass(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = UpdatePassSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.PassesService/DeletePass" => {
                    #[allow(non_camel_case_types)]
                    struct DeletePassSvc<T: PassesService>(pub Arc<T>);
                    impl<
                        T: PassesService,
                    > tonic::server::UnaryService<super::DeletePassRequest>
                    for DeletePassSvc<T> {
                        type Response = super::DeletePassResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::DeletePassRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as PassesService>::delete_pass(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = DeletePassSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.PassesService/ListPassRequirements" => {
                    #[allow(non_camel_case_types)]
                    struct ListPassRequirementsSvc<T: PassesService>(pub Arc<T>);
                    impl<
                        T: PassesService,
                    > tonic::server::UnaryService<super::ListPassRequirementsRequest>
                    for ListPassRequirementsSvc<T> {
                        type Response = super::ListPassRequirementsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListPassRequirementsRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as PassesService>::list_pass_requirements(
                                        &inner,
                                        request,
                                    )
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ListPassRequirementsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.PassesService/CreatePassRequirement" => {
                    #[allow(non_camel_case_types)]
                    struct CreatePassRequirementSvc<T: PassesService>(pub Arc<T>);
                    impl<
                        T: PassesService,
                    > tonic::server::UnaryService<super::CreatePassRequirementRequest>
                    for CreatePassRequirementSvc<T> {
                        type Response = super::CreatePassRequirementResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CreatePassRequirementRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as PassesService>::create_pass_requirement(
                                        &inner,
                                        request,
                                    )
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = CreatePassRequirementSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.PassesService/DeletePassRequirement" => {
                    #[allow(non_camel_case_types)]
                    struct DeletePassRequirementSvc<T: PassesService>(pub Arc<T>);
                    impl<
                        T: PassesService,
                    > tonic::server::UnaryService<super::DeletePassRequirementRequest>
                    for DeletePassRequirementSvc<T> {
                        type Response = super::DeletePassRequirementResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::DeletePassRequirementRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as PassesService>::delete_pass_requirement(
                                        &inner,
                                        request,
                                    )
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = DeletePassRequirementSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.PassesService/CanSubscribeToPass" => {
                    #[allow(non_camel_case_types)]
                    struct CanSubscribeToPassSvc<T: PassesService>(pub Arc<T>);
                    impl<
                        T: PassesService,
                    > tonic::server::UnaryService<super::CanSubscribeToPassRequest>
                    for CanSubscribeToPassSvc<T> {
                        type Response = super::CanSubscribeToPassResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CanSubscribeToPassRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as PassesService>::can_subscribe_to_pass(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = CanSubscribeToPassSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", tonic::Code::Unimplemented as i32)
                                .header(
                                    http::header::CONTENT_TYPE,
                                    tonic::metadata::GRPC_CONTENT_TYPE,
                                )
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: PassesService> Clone for PassesServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: PassesService> tonic::server::NamedService for PassesServiceServer<T> {
        const NAME: &'static str = "tap.ads.v1.PassesService";
    }
}
/// Generated client implementations.
pub mod quotas_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct QuotasServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl QuotasServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> QuotasServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> QuotasServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            QuotasServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn list_quotas(
            &mut self,
            request: impl tonic::IntoRequest<super::ListQuotasRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListQuotasResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.QuotasService/ListQuotas",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.QuotasService", "ListQuotas"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_quota(
            &mut self,
            request: impl tonic::IntoRequest<super::GetQuotaRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetQuotaResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.QuotasService/GetQuota",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.QuotasService", "GetQuota"));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod quotas_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with QuotasServiceServer.
    #[async_trait]
    pub trait QuotasService: Send + Sync + 'static {
        async fn list_quotas(
            &self,
            request: tonic::Request<super::ListQuotasRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListQuotasResponse>,
            tonic::Status,
        >;
        async fn get_quota(
            &self,
            request: tonic::Request<super::GetQuotaRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetQuotaResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct QuotasServiceServer<T: QuotasService> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T: QuotasService> QuotasServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for QuotasServiceServer<T>
    where
        T: QuotasService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/tap.ads.v1.QuotasService/ListQuotas" => {
                    #[allow(non_camel_case_types)]
                    struct ListQuotasSvc<T: QuotasService>(pub Arc<T>);
                    impl<
                        T: QuotasService,
                    > tonic::server::UnaryService<super::ListQuotasRequest>
                    for ListQuotasSvc<T> {
                        type Response = super::ListQuotasResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListQuotasRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as QuotasService>::list_quotas(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ListQuotasSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.QuotasService/GetQuota" => {
                    #[allow(non_camel_case_types)]
                    struct GetQuotaSvc<T: QuotasService>(pub Arc<T>);
                    impl<
                        T: QuotasService,
                    > tonic::server::UnaryService<super::GetQuotaRequest>
                    for GetQuotaSvc<T> {
                        type Response = super::GetQuotaResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GetQuotaRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as QuotasService>::get_quota(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = GetQuotaSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", tonic::Code::Unimplemented as i32)
                                .header(
                                    http::header::CONTENT_TYPE,
                                    tonic::metadata::GRPC_CONTENT_TYPE,
                                )
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: QuotasService> Clone for QuotasServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: QuotasService> tonic::server::NamedService for QuotasServiceServer<T> {
        const NAME: &'static str = "tap.ads.v1.QuotasService";
    }
}
/// Generated client implementations.
pub mod ads_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct AdsServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl AdsServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> AdsServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> AdsServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            AdsServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn list_ads(
            &mut self,
            request: impl tonic::IntoRequest<super::ListAdsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListAdsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.AdsService/ListAds",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.AdsService", "ListAds"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_ad(
            &mut self,
            request: impl tonic::IntoRequest<super::GetAdRequest>,
        ) -> std::result::Result<tonic::Response<super::GetAdResponse>, tonic::Status> {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.AdsService/GetAd",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.AdsService", "GetAd"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn create_ad(
            &mut self,
            request: impl tonic::IntoRequest<super::CreateAdRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateAdResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.AdsService/CreateAd",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.AdsService", "CreateAd"));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod ads_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with AdsServiceServer.
    #[async_trait]
    pub trait AdsService: Send + Sync + 'static {
        async fn list_ads(
            &self,
            request: tonic::Request<super::ListAdsRequest>,
        ) -> std::result::Result<tonic::Response<super::ListAdsResponse>, tonic::Status>;
        async fn get_ad(
            &self,
            request: tonic::Request<super::GetAdRequest>,
        ) -> std::result::Result<tonic::Response<super::GetAdResponse>, tonic::Status>;
        async fn create_ad(
            &self,
            request: tonic::Request<super::CreateAdRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateAdResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct AdsServiceServer<T: AdsService> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T: AdsService> AdsServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for AdsServiceServer<T>
    where
        T: AdsService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/tap.ads.v1.AdsService/ListAds" => {
                    #[allow(non_camel_case_types)]
                    struct ListAdsSvc<T: AdsService>(pub Arc<T>);
                    impl<
                        T: AdsService,
                    > tonic::server::UnaryService<super::ListAdsRequest>
                    for ListAdsSvc<T> {
                        type Response = super::ListAdsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListAdsRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as AdsService>::list_ads(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ListAdsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.AdsService/GetAd" => {
                    #[allow(non_camel_case_types)]
                    struct GetAdSvc<T: AdsService>(pub Arc<T>);
                    impl<T: AdsService> tonic::server::UnaryService<super::GetAdRequest>
                    for GetAdSvc<T> {
                        type Response = super::GetAdResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GetAdRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as AdsService>::get_ad(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = GetAdSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.AdsService/CreateAd" => {
                    #[allow(non_camel_case_types)]
                    struct CreateAdSvc<T: AdsService>(pub Arc<T>);
                    impl<
                        T: AdsService,
                    > tonic::server::UnaryService<super::CreateAdRequest>
                    for CreateAdSvc<T> {
                        type Response = super::CreateAdResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CreateAdRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as AdsService>::create_ad(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = CreateAdSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", tonic::Code::Unimplemented as i32)
                                .header(
                                    http::header::CONTENT_TYPE,
                                    tonic::metadata::GRPC_CONTENT_TYPE,
                                )
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: AdsService> Clone for AdsServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: AdsService> tonic::server::NamedService for AdsServiceServer<T> {
        const NAME: &'static str = "tap.ads.v1.AdsService";
    }
}
/// Generated client implementations.
pub mod documents_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct DocumentsServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl DocumentsServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> DocumentsServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> DocumentsServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            DocumentsServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn list_documents(
            &mut self,
            request: impl tonic::IntoRequest<super::ListDocumentsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListDocumentsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.DocumentsService/ListDocuments",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.DocumentsService", "ListDocuments"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_document(
            &mut self,
            request: impl tonic::IntoRequest<super::GetDocumentRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetDocumentResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.DocumentsService/GetDocument",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.DocumentsService", "GetDocument"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn create_document(
            &mut self,
            request: impl tonic::IntoRequest<super::CreateDocumentRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateDocumentResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.DocumentsService/CreateDocument",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.DocumentsService", "CreateDocument"),
                );
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod documents_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with DocumentsServiceServer.
    #[async_trait]
    pub trait DocumentsService: Send + Sync + 'static {
        async fn list_documents(
            &self,
            request: tonic::Request<super::ListDocumentsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListDocumentsResponse>,
            tonic::Status,
        >;
        async fn get_document(
            &self,
            request: tonic::Request<super::GetDocumentRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetDocumentResponse>,
            tonic::Status,
        >;
        async fn create_document(
            &self,
            request: tonic::Request<super::CreateDocumentRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateDocumentResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct DocumentsServiceServer<T: DocumentsService> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T: DocumentsService> DocumentsServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for DocumentsServiceServer<T>
    where
        T: DocumentsService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/tap.ads.v1.DocumentsService/ListDocuments" => {
                    #[allow(non_camel_case_types)]
                    struct ListDocumentsSvc<T: DocumentsService>(pub Arc<T>);
                    impl<
                        T: DocumentsService,
                    > tonic::server::UnaryService<super::ListDocumentsRequest>
                    for ListDocumentsSvc<T> {
                        type Response = super::ListDocumentsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListDocumentsRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as DocumentsService>::list_documents(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ListDocumentsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.DocumentsService/GetDocument" => {
                    #[allow(non_camel_case_types)]
                    struct GetDocumentSvc<T: DocumentsService>(pub Arc<T>);
                    impl<
                        T: DocumentsService,
                    > tonic::server::UnaryService<super::GetDocumentRequest>
                    for GetDocumentSvc<T> {
                        type Response = super::GetDocumentResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GetDocumentRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as DocumentsService>::get_document(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = GetDocumentSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.DocumentsService/CreateDocument" => {
                    #[allow(non_camel_case_types)]
                    struct CreateDocumentSvc<T: DocumentsService>(pub Arc<T>);
                    impl<
                        T: DocumentsService,
                    > tonic::server::UnaryService<super::CreateDocumentRequest>
                    for CreateDocumentSvc<T> {
                        type Response = super::CreateDocumentResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CreateDocumentRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as DocumentsService>::create_document(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = CreateDocumentSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", tonic::Code::Unimplemented as i32)
                                .header(
                                    http::header::CONTENT_TYPE,
                                    tonic::metadata::GRPC_CONTENT_TYPE,
                                )
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: DocumentsService> Clone for DocumentsServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: DocumentsService> tonic::server::NamedService for DocumentsServiceServer<T> {
        const NAME: &'static str = "tap.ads.v1.DocumentsService";
    }
}
/// Generated client implementations.
pub mod listings_service_client {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    use tonic::codegen::http::Uri;
    #[derive(Debug, Clone)]
    pub struct ListingsServiceClient<T> {
        inner: tonic::client::Grpc<T>,
    }
    impl ListingsServiceClient<tonic::transport::Channel> {
        /// Attempt to create a new client by connecting to a given endpoint.
        pub async fn connect<D>(dst: D) -> Result<Self, tonic::transport::Error>
        where
            D: TryInto<tonic::transport::Endpoint>,
            D::Error: Into<StdError>,
        {
            let conn = tonic::transport::Endpoint::new(dst)?.connect().await?;
            Ok(Self::new(conn))
        }
    }
    impl<T> ListingsServiceClient<T>
    where
        T: tonic::client::GrpcService<tonic::body::BoxBody>,
        T::Error: Into<StdError>,
        T::ResponseBody: Body<Data = Bytes> + Send + 'static,
        <T::ResponseBody as Body>::Error: Into<StdError> + Send,
    {
        pub fn new(inner: T) -> Self {
            let inner = tonic::client::Grpc::new(inner);
            Self { inner }
        }
        pub fn with_origin(inner: T, origin: Uri) -> Self {
            let inner = tonic::client::Grpc::with_origin(inner, origin);
            Self { inner }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> ListingsServiceClient<InterceptedService<T, F>>
        where
            F: tonic::service::Interceptor,
            T::ResponseBody: Default,
            T: tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
                Response = http::Response<
                    <T as tonic::client::GrpcService<tonic::body::BoxBody>>::ResponseBody,
                >,
            >,
            <T as tonic::codegen::Service<
                http::Request<tonic::body::BoxBody>,
            >>::Error: Into<StdError> + Send + Sync,
        {
            ListingsServiceClient::new(InterceptedService::new(inner, interceptor))
        }
        /// Compress requests with the given encoding.
        ///
        /// This requires the server to support it otherwise it might respond with an
        /// error.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.send_compressed(encoding);
            self
        }
        /// Enable decompressing responses.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.inner = self.inner.accept_compressed(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_decoding_message_size(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.inner = self.inner.max_encoding_message_size(limit);
            self
        }
        pub async fn list_listings(
            &mut self,
            request: impl tonic::IntoRequest<super::ListListingsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListListingsResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.ListingsService/ListListings",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.ListingsService", "ListListings"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn get_listing(
            &mut self,
            request: impl tonic::IntoRequest<super::GetListingRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetListingResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.ListingsService/GetListing",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.ListingsService", "GetListing"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn create_listing(
            &mut self,
            request: impl tonic::IntoRequest<super::CreateListingRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateListingResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.ListingsService/CreateListing",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.ListingsService", "CreateListing"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn delete_listing(
            &mut self,
            request: impl tonic::IntoRequest<super::DeleteListingRequest>,
        ) -> std::result::Result<
            tonic::Response<super::DeleteListingResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.ListingsService/DeleteListing",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.ListingsService", "DeleteListing"));
            self.inner.unary(req, path, codec).await
        }
        pub async fn purchase_listing(
            &mut self,
            request: impl tonic::IntoRequest<super::PurchaseListingRequest>,
        ) -> std::result::Result<
            tonic::Response<super::PurchaseListingResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.ListingsService/PurchaseListing",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(
                    GrpcMethod::new("tap.ads.v1.ListingsService", "PurchaseListing"),
                );
            self.inner.unary(req, path, codec).await
        }
        pub async fn list_purchases(
            &mut self,
            request: impl tonic::IntoRequest<super::ListPurchasesRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListPurchasesResponse>,
            tonic::Status,
        > {
            self.inner
                .ready()
                .await
                .map_err(|e| {
                    tonic::Status::new(
                        tonic::Code::Unknown,
                        format!("Service was not ready: {}", e.into()),
                    )
                })?;
            let codec = tonic::codec::ProstCodec::default();
            let path = http::uri::PathAndQuery::from_static(
                "/tap.ads.v1.ListingsService/ListPurchases",
            );
            let mut req = request.into_request();
            req.extensions_mut()
                .insert(GrpcMethod::new("tap.ads.v1.ListingsService", "ListPurchases"));
            self.inner.unary(req, path, codec).await
        }
    }
}
/// Generated server implementations.
pub mod listings_service_server {
    #![allow(unused_variables, dead_code, missing_docs, clippy::let_unit_value)]
    use tonic::codegen::*;
    /// Generated trait containing gRPC methods that should be implemented for use with ListingsServiceServer.
    #[async_trait]
    pub trait ListingsService: Send + Sync + 'static {
        async fn list_listings(
            &self,
            request: tonic::Request<super::ListListingsRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListListingsResponse>,
            tonic::Status,
        >;
        async fn get_listing(
            &self,
            request: tonic::Request<super::GetListingRequest>,
        ) -> std::result::Result<
            tonic::Response<super::GetListingResponse>,
            tonic::Status,
        >;
        async fn create_listing(
            &self,
            request: tonic::Request<super::CreateListingRequest>,
        ) -> std::result::Result<
            tonic::Response<super::CreateListingResponse>,
            tonic::Status,
        >;
        async fn delete_listing(
            &self,
            request: tonic::Request<super::DeleteListingRequest>,
        ) -> std::result::Result<
            tonic::Response<super::DeleteListingResponse>,
            tonic::Status,
        >;
        async fn purchase_listing(
            &self,
            request: tonic::Request<super::PurchaseListingRequest>,
        ) -> std::result::Result<
            tonic::Response<super::PurchaseListingResponse>,
            tonic::Status,
        >;
        async fn list_purchases(
            &self,
            request: tonic::Request<super::ListPurchasesRequest>,
        ) -> std::result::Result<
            tonic::Response<super::ListPurchasesResponse>,
            tonic::Status,
        >;
    }
    #[derive(Debug)]
    pub struct ListingsServiceServer<T: ListingsService> {
        inner: Arc<T>,
        accept_compression_encodings: EnabledCompressionEncodings,
        send_compression_encodings: EnabledCompressionEncodings,
        max_decoding_message_size: Option<usize>,
        max_encoding_message_size: Option<usize>,
    }
    impl<T: ListingsService> ListingsServiceServer<T> {
        pub fn new(inner: T) -> Self {
            Self::from_arc(Arc::new(inner))
        }
        pub fn from_arc(inner: Arc<T>) -> Self {
            Self {
                inner,
                accept_compression_encodings: Default::default(),
                send_compression_encodings: Default::default(),
                max_decoding_message_size: None,
                max_encoding_message_size: None,
            }
        }
        pub fn with_interceptor<F>(
            inner: T,
            interceptor: F,
        ) -> InterceptedService<Self, F>
        where
            F: tonic::service::Interceptor,
        {
            InterceptedService::new(Self::new(inner), interceptor)
        }
        /// Enable decompressing requests with the given encoding.
        #[must_use]
        pub fn accept_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.accept_compression_encodings.enable(encoding);
            self
        }
        /// Compress responses with the given encoding, if the client supports it.
        #[must_use]
        pub fn send_compressed(mut self, encoding: CompressionEncoding) -> Self {
            self.send_compression_encodings.enable(encoding);
            self
        }
        /// Limits the maximum size of a decoded message.
        ///
        /// Default: `4MB`
        #[must_use]
        pub fn max_decoding_message_size(mut self, limit: usize) -> Self {
            self.max_decoding_message_size = Some(limit);
            self
        }
        /// Limits the maximum size of an encoded message.
        ///
        /// Default: `usize::MAX`
        #[must_use]
        pub fn max_encoding_message_size(mut self, limit: usize) -> Self {
            self.max_encoding_message_size = Some(limit);
            self
        }
    }
    impl<T, B> tonic::codegen::Service<http::Request<B>> for ListingsServiceServer<T>
    where
        T: ListingsService,
        B: Body + Send + 'static,
        B::Error: Into<StdError> + Send + 'static,
    {
        type Response = http::Response<tonic::body::BoxBody>;
        type Error = std::convert::Infallible;
        type Future = BoxFuture<Self::Response, Self::Error>;
        fn poll_ready(
            &mut self,
            _cx: &mut Context<'_>,
        ) -> Poll<std::result::Result<(), Self::Error>> {
            Poll::Ready(Ok(()))
        }
        fn call(&mut self, req: http::Request<B>) -> Self::Future {
            match req.uri().path() {
                "/tap.ads.v1.ListingsService/ListListings" => {
                    #[allow(non_camel_case_types)]
                    struct ListListingsSvc<T: ListingsService>(pub Arc<T>);
                    impl<
                        T: ListingsService,
                    > tonic::server::UnaryService<super::ListListingsRequest>
                    for ListListingsSvc<T> {
                        type Response = super::ListListingsResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListListingsRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ListingsService>::list_listings(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ListListingsSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.ListingsService/GetListing" => {
                    #[allow(non_camel_case_types)]
                    struct GetListingSvc<T: ListingsService>(pub Arc<T>);
                    impl<
                        T: ListingsService,
                    > tonic::server::UnaryService<super::GetListingRequest>
                    for GetListingSvc<T> {
                        type Response = super::GetListingResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::GetListingRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ListingsService>::get_listing(&inner, request).await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = GetListingSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.ListingsService/CreateListing" => {
                    #[allow(non_camel_case_types)]
                    struct CreateListingSvc<T: ListingsService>(pub Arc<T>);
                    impl<
                        T: ListingsService,
                    > tonic::server::UnaryService<super::CreateListingRequest>
                    for CreateListingSvc<T> {
                        type Response = super::CreateListingResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::CreateListingRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ListingsService>::create_listing(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = CreateListingSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.ListingsService/DeleteListing" => {
                    #[allow(non_camel_case_types)]
                    struct DeleteListingSvc<T: ListingsService>(pub Arc<T>);
                    impl<
                        T: ListingsService,
                    > tonic::server::UnaryService<super::DeleteListingRequest>
                    for DeleteListingSvc<T> {
                        type Response = super::DeleteListingResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::DeleteListingRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ListingsService>::delete_listing(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = DeleteListingSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.ListingsService/PurchaseListing" => {
                    #[allow(non_camel_case_types)]
                    struct PurchaseListingSvc<T: ListingsService>(pub Arc<T>);
                    impl<
                        T: ListingsService,
                    > tonic::server::UnaryService<super::PurchaseListingRequest>
                    for PurchaseListingSvc<T> {
                        type Response = super::PurchaseListingResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::PurchaseListingRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ListingsService>::purchase_listing(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = PurchaseListingSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                "/tap.ads.v1.ListingsService/ListPurchases" => {
                    #[allow(non_camel_case_types)]
                    struct ListPurchasesSvc<T: ListingsService>(pub Arc<T>);
                    impl<
                        T: ListingsService,
                    > tonic::server::UnaryService<super::ListPurchasesRequest>
                    for ListPurchasesSvc<T> {
                        type Response = super::ListPurchasesResponse;
                        type Future = BoxFuture<
                            tonic::Response<Self::Response>,
                            tonic::Status,
                        >;
                        fn call(
                            &mut self,
                            request: tonic::Request<super::ListPurchasesRequest>,
                        ) -> Self::Future {
                            let inner = Arc::clone(&self.0);
                            let fut = async move {
                                <T as ListingsService>::list_purchases(&inner, request)
                                    .await
                            };
                            Box::pin(fut)
                        }
                    }
                    let accept_compression_encodings = self.accept_compression_encodings;
                    let send_compression_encodings = self.send_compression_encodings;
                    let max_decoding_message_size = self.max_decoding_message_size;
                    let max_encoding_message_size = self.max_encoding_message_size;
                    let inner = self.inner.clone();
                    let fut = async move {
                        let method = ListPurchasesSvc(inner);
                        let codec = tonic::codec::ProstCodec::default();
                        let mut grpc = tonic::server::Grpc::new(codec)
                            .apply_compression_config(
                                accept_compression_encodings,
                                send_compression_encodings,
                            )
                            .apply_max_message_size_config(
                                max_decoding_message_size,
                                max_encoding_message_size,
                            );
                        let res = grpc.unary(method, req).await;
                        Ok(res)
                    };
                    Box::pin(fut)
                }
                _ => {
                    Box::pin(async move {
                        Ok(
                            http::Response::builder()
                                .status(200)
                                .header("grpc-status", tonic::Code::Unimplemented as i32)
                                .header(
                                    http::header::CONTENT_TYPE,
                                    tonic::metadata::GRPC_CONTENT_TYPE,
                                )
                                .body(empty_body())
                                .unwrap(),
                        )
                    })
                }
            }
        }
    }
    impl<T: ListingsService> Clone for ListingsServiceServer<T> {
        fn clone(&self) -> Self {
            let inner = self.inner.clone();
            Self {
                inner,
                accept_compression_encodings: self.accept_compression_encodings,
                send_compression_encodings: self.send_compression_encodings,
                max_decoding_message_size: self.max_decoding_message_size,
                max_encoding_message_size: self.max_encoding_message_size,
            }
        }
    }
    impl<T: ListingsService> tonic::server::NamedService for ListingsServiceServer<T> {
        const NAME: &'static str = "tap.ads.v1.ListingsService";
    }
}
