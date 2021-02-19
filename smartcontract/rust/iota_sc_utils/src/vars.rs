//use wasmlib::*;

// /// Tries to get &str parameter. Panics if it can't find it.
// pub fn must_get_string<T:ScBaseContext>(var_name : &str, ctx: &T) -> String {
//     let var = ctx.state().get_string(var_name);
//     ctx.require(var.exists(), "string parameter not found");
//     return var.value();
// }

// /// Tries to get &str parameter. Returns empty &str if it can't find it.
// pub fn get_string(var_name : &str, ctx: &ScFuncContext) -> String {
//     let var = ctx.state().get_string(var_name);
//     if var.exists() {
//         return var.value();
//     }
//     return String::from("");
// }

// /// Stores "value" in a variable "var_name"
// pub fn set_string(var_name : &str, value : &str, ctx: &ScFuncContext) {
//     ctx.state().get_string(var_name).set_value(value);
// }
