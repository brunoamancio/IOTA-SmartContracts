use wasmlib::*;

/// Tries to get &str parameter. Panics if it can't find it.
pub fn must_get_string<TContext : ScBaseContext>(param_name : &str, ctx: &TContext) -> String {
    let param = ctx.params().get_string(param_name);
    ctx.require(param.exists(), "string parameter not found");
    return param.value();
}

/// Tries to get &str parameter. Returns empty &str if it can't find it.
pub fn get_string<TContext : ScBaseContext>(param_name : &str, ctx: &TContext) -> String {
    let param = ctx.params().get_string(param_name);
    if param.exists() {
        return param.value();
    }
    return String::from("");
}

/// Tries to get i64 parameter. Panics if it can't find it.
pub fn must_get_int64<TContext : ScBaseContext>(param_name : &str, ctx: &TContext) -> i64 {
    let param = ctx.params().get_int(param_name);
    ctx.require(param.exists(), "string parameter not found");
    return param.value();
}

/// Tries to get i64 parameter. Returns empty &str if it can't find it.
pub fn get_int64<TContext : ScBaseContext>(param_name : &str, ctx: &TContext) -> i64 {
    let param = ctx.params().get_int(param_name);
    if param.exists() {
        return param.value();
    }
    return  0;
}

