use wasmlib::*;
use iota_sc_utils::params;

// ################################################################################################## //
// These code samples are not executed anywhere. It just show samples of how you can call in your SC. //
// ################################################################################################## //

pub const MY_PARAMETER: &str = "my_parameter";

pub fn how_to_get_string_params_not_panic_if_not_found(ctx: &ScFuncContext) -> String {
    let string_parameter1 = params::get_string(MY_PARAMETER, ctx);
    return string_parameter1;
}

pub fn how_to_get_string_params_panic_if_not_found(ctx: &ScFuncContext) -> String {
    let string_parameter2 = params::must_get_string(MY_PARAMETER, ctx);
    return string_parameter2;
}

pub fn how_to_get_int64_params_not_panic_if_not_found(ctx: &ScFuncContext) -> i64 {
    let int64_parameter1 = params::get_int64(MY_PARAMETER, ctx);
    return int64_parameter1;
}

pub fn how_to_get_int64_params_panic_if_not_found(ctx: &ScFuncContext) -> i64 {
    let int64_parameter2 = params::must_get_int64(MY_PARAMETER, ctx);
    return int64_parameter2;
}

// pub fn how_to_get_int64_params_panic_if_not_found(ctx: &ScFuncContext) -> String{
//         // Reads a string parameter passed to SC function call. It is empty if not found.
//         pub const MY_PARAMETER: &str = "my_parameter";
//         let string_parameter1 : String = params::get_string(MY_PARAMETER, ctx);
//         if !string_parameter1.is_empty() {
//             ctx.log(&string_parameter1);
//         }
    
//         // Sets value of "MY_VARIABLE" to "my dummy value".
//         const MY_VARIABLE : &str = "my_var";
//         vars::set_string(MY_VARIABLE, "my dummy value", ctx);
    
//         // Reads the value of "MY_VARIABLE". It is empty if not set.
//         let variable = vars::get_string(MY_VARIABLE, ctx);
//         if !variable.is_empty() {
//             ctx.log(&string_parameter1);
//         }
// }