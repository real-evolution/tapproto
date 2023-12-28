mod compat;
mod gen;

#[doc(inline)]
pub use gen::tap::*;

pub mod thirdparty {
    #[doc(inline)]
    pub use super::gen::google;

    #[doc(inline)]
    pub use super::gen::buf;
}
