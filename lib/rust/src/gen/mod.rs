// @generated
pub mod buf {
    #[cfg(feature = "buf-validate")]
    // @@protoc_insertion_point(attribute:buf.validate)
    pub mod validate {
        include!("buf.validate.rs");
        // @@protoc_insertion_point(buf.validate)
        #[cfg(feature = "buf-validate-priv")]
        // @@protoc_insertion_point(attribute:buf.validate.priv)
        pub mod r#priv {
            include!("buf.validate.priv.rs");
            // @@protoc_insertion_point(buf.validate.priv)
        }
    }
}
pub mod google {
    #[cfg(feature = "google-api")]
    // @@protoc_insertion_point(attribute:google.api)
    pub mod api {
        include!("google.api.rs");
        // @@protoc_insertion_point(google.api)
    }
    #[cfg(feature = "google-protobuf")]
    // @@protoc_insertion_point(attribute:google.protobuf)
    pub mod protobuf {
        include!("google.protobuf.rs");
        // @@protoc_insertion_point(google.protobuf)
    }
    #[cfg(feature = "google-type")]
    // @@protoc_insertion_point(attribute:google.type)
    pub mod r#type {
        include!("google.type.rs");
        // @@protoc_insertion_point(google.type)
    }
}
pub mod tap {
    pub mod ads {
        #[cfg(feature = "tap-ads-v1")]
        // @@protoc_insertion_point(attribute:tap.ads.v1)
        pub mod v1 {
            include!("tap.ads.v1.rs");
            // @@protoc_insertion_point(tap.ads.v1)
        }
    }
    pub mod campaign {
        #[cfg(feature = "tap-campaign-v1")]
        // @@protoc_insertion_point(attribute:tap.campaign.v1)
        pub mod v1 {
            include!("tap.campaign.v1.rs");
            // @@protoc_insertion_point(tap.campaign.v1)
        }
    }
    pub mod catalog {
        #[cfg(feature = "tap-catalog-v1")]
        // @@protoc_insertion_point(attribute:tap.catalog.v1)
        pub mod v1 {
            include!("tap.catalog.v1.rs");
            // @@protoc_insertion_point(tap.catalog.v1)
        }
    }
    pub mod points {
        #[cfg(feature = "tap-points-v1")]
        // @@protoc_insertion_point(attribute:tap.points.v1)
        pub mod v1 {
            include!("tap.points.v1.rs");
            // @@protoc_insertion_point(tap.points.v1)
        }
    }
    pub mod sub {
        #[cfg(feature = "tap-sub-v1")]
        // @@protoc_insertion_point(attribute:tap.sub.v1)
        pub mod v1 {
            include!("tap.sub.v1.rs");
            // @@protoc_insertion_point(tap.sub.v1)
        }
    }
    pub mod types {
        #[cfg(feature = "tap-types-v1")]
        // @@protoc_insertion_point(attribute:tap.types.v1)
        pub mod v1 {
            include!("tap.types.v1.rs");
            // @@protoc_insertion_point(tap.types.v1)
        }
    }
    pub mod verify {
        #[cfg(feature = "tap-verify-v1")]
        // @@protoc_insertion_point(attribute:tap.verify.v1)
        pub mod v1 {
            include!("tap.verify.v1.rs");
            // @@protoc_insertion_point(tap.verify.v1)
        }
    }
}