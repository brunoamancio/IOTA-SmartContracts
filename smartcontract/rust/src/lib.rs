use wasmlib::*;
use iota_sc_utils::*;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();

    // Functions
    exports.add_func("my_sc_function", my_sc_function);
    exports.add_func("contract_creator_only_function", contract_creator_only_function);
    exports.add_func("chain_owner_only_function", chain_owner_only_function);

    // Views
    exports.add_view("my_sc_view", my_sc_view);
}

// Anyone can call this SC-Function
fn my_sc_function(ctx: &ScFuncContext) {
    // Logs a text
    ctx.log("my_sc_function");

    // Reads argument called "my_param" passed to SC function. Empty if not found.
    const MY_PARAM : &str = "my_param";
    let param_value = params::get_string(MY_PARAM, ctx);
    if !param_value.is_empty() {
        ctx.log(&param_value);
    }
}

// Only the contract creator can call this SC-Function
fn contract_creator_only_function(ctx: &ScFuncContext) {
    access::caller_must_be_contract_creator(ctx);
    ctx.log("Caller is the contract creator =)");
}

// Only the chain owner can call this SC-Function
fn chain_owner_only_function(ctx: &ScFuncContext){
    access::caller_must_be_chain_owner(ctx);
    ctx.log("Caller is the chain owner =)");
}

// Public view
fn my_sc_view(ctx: &ScViewContext) {
    ctx.log("my_sc_view");
}