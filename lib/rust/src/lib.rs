mod compat;
mod gen;

#[doc(inline)]
pub use gen::tap::*;

pub mod thirdparty {
    #[doc(inline)]
    pub use super::gen::google;
}
