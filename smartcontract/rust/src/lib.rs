use wasmlib::*;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_call("my_sc_request", my_sc_request);
    exports.add_view("my_sc_view", my_sc_view);
}

fn my_sc_request(ctx: &ScCallContext) {
    ctx.log("my_sc_request");
}

fn my_sc_view(ctx: &ScViewContext) {
    ctx.log("my_sc_view");
}