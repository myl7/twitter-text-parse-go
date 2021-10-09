use libc::c_char;
use std::ffi::CStr;
use std::ptr::null;
use twitter_text_config::Configuration;

#[repr(C)]
pub struct TwTextParseRes {
    weighted_length: i32,
    permillage: i32,
    is_valid: c_char,
    display_text_range_start: i32,
    display_text_range_end: i32,
    valid_text_range_start: i32,
    valid_text_range_end: i32,
}

#[no_mangle]
pub extern "C" fn tw_text_parse(s: *const c_char) -> *const TwTextParseRes {
    if s.is_null() {
        return null();
    }
    let text = unsafe { CStr::from_ptr(s) }.to_str().unwrap();
    let config = Configuration::default();
    let raw_res = twitter_text::parse(text, &config, true);
    let res = Box::new(TwTextParseRes {
        weighted_length: raw_res.weighted_length,
        permillage: raw_res.permillage,
        is_valid: raw_res.is_valid as c_char,
        display_text_range_start: raw_res.display_text_range.start(),
        display_text_range_end: raw_res.display_text_range.end(),
        valid_text_range_start: raw_res.valid_text_range.start(),
        valid_text_range_end: raw_res.valid_text_range.end(),
    });
    return Box::into_raw(res);
}
