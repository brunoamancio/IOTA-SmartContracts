use wasmlib::*;

// -------------------------------------------------------------------------------------------------------------------- //
//  Feel free to add any functions related to params here. Don't forget to make a pull request to the template repo. :) //
// -------------------------------------------------------------------------------------------------------------------- //

/// Tries to set &str parameter. Does nothing if it can't find it.
pub fn set_string<TContext : ScBaseContext>(param_name : &str, param_value : &str, ctx: &TContext) {
    ctx.results().get_string(param_name).set_value(param_value);
}

/// Tries to set i64 parameter. Does nothing if it can't find it.
pub fn set_i64<TContext : ScBaseContext>(param_name : &str, param_value : i64, ctx: &TContext) {
    ctx.results().get_int64(param_name).set_value(param_value);
}

/// Tries to set bytes parameter. Does nothing if it can't find it.
pub fn set_bytes<TContext : ScBaseContext>(param_name : &str, param_value : &[u8], ctx: &TContext) {
    ctx.results().get_bytes(param_name).set_value(param_value);
}

/// Sets bool parameter. 
pub fn set_bool<TContext : ScBaseContext>(param_name : &str, param_value : bool, ctx: &TContext) {
    let bool_value = to_u8(param_value);
    ctx.results().get_bytes(param_name).set_value(&bool_value);
}

/// converts a boolean value into a vector of u8
fn to_u8(param_value : bool) -> Vec<u8> {
    let mut value : Vec<u8> = Vec::new();
    if param_value {
        value.push(1)
    } else {
        value.push(0)
    }
    value
}