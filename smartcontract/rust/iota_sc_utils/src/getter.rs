use wasmlib::*;

pub const STATE : &str = "state";
pub const PARAMS : &str = "params";

macro_rules! add_impl_pub_fns {

    ($source:tt, $must_get_func_name:ident, $get_func_name:ident, $return_type:ty) => {
        /// Tries to get a variable. Panics if it can't find it.
        pub fn $must_get_func_name<TGetter:Getter>(variable_name : &str, context : &TGetter) -> $return_type {
            context.$must_get_func_name($source, variable_name)
        }

        /// Tries to get a variable. Returns default value if it can't find it.
        pub fn $get_func_name<TGetter:Getter>(variable_name : &str, context : &TGetter) -> $return_type {
            context.$get_func_name($source, variable_name)
        }
    };
}

macro_rules! add_all_getter_fns {
    ($must_get_func_name:ident, $get_func_name:ident, $return_type:ty) => {
        /// Tries to get a variable. Panics if it can't find it.
        fn $must_get_func_name(&self, source : &str, variable_name : &str) -> $return_type;
        /// Tries to get a variable. Returns default value if it can't find it.
        fn $get_func_name(&self, source : &str, variable_name : &str) -> $return_type;
    };

    () => {
        // Primitive types
        add_all_getter_fns!(must_get_string, get_string, String);
        add_all_getter_fns!(must_get_int64, get_int64, i64);
        add_all_getter_fns!(must_get_bytes, get_bytes, Vec<u8>);

        // ISCP types
        add_all_getter_fns!(must_get_agent_id, get_agent_id, ScAgentId);
        add_all_getter_fns!(must_get_address, get_address, ScAddress);
        add_all_getter_fns!(must_get_request_id, get_request_id, ScRequestId);
        add_all_getter_fns!(must_get_hname, get_hname, ScHname);
        add_all_getter_fns!(must_get_hash, get_hash, ScHash);
        add_all_getter_fns!(must_get_contract_id, get_contract_id, ScContractId);
        add_all_getter_fns!(must_get_color, get_color, ScColor);
        add_all_getter_fns!(must_get_chain_id, get_chain_id, ScChainId);
    };
}

pub trait Getter {
    add_all_getter_fns!();
}

macro_rules! add_impl_getters {
    ($must_get_func_name:ident, $get_func_name:ident, $return_type:ty) => {
        
        /// Tries to get a variable. Panics if it can't find it.
        fn $must_get_func_name(&self, source : &str, variable_name : &str) -> $return_type {
            match source {
                crate::getter::STATE =>  {
                    let variable = self.state().$get_func_name(variable_name);
                    crate::getter::require_if_needed(self, variable.exists(), &format!("variable {} not found", variable_name));
                    variable.value()
                },
                crate::getter::PARAMS => {
                    let param = self.params().$get_func_name(variable_name);
                    crate::getter::require_if_needed(self, param.exists(), &format!("parameter {} not found", variable_name));
                    param.value()
                },
                _ => panic!(format!("Source {} not implemented", source)),
            }
        }

        /// Tries to get a variable. Returns default value if it can't find it.
        fn $get_func_name(&self, source : &str, variable_name : &str) -> $return_type {
            match source {
                crate::getter::STATE =>  {
                    self.state().$get_func_name(variable_name).value()
                },
                crate::getter::PARAMS => {
                    self.params().$get_func_name(variable_name).value()
                },
                _ => panic!(format!("Source {} not implemented", source)),
            }
        }
    };

    ($context:ty) => {
        impl Getter for $context {
            // Primitive types
            add_impl_getters!(must_get_string, get_string, String);
            add_impl_getters!(must_get_int64, get_int64, i64);
            add_impl_getters!(must_get_bytes, get_bytes, Vec<u8>);
        
            // ISCP types
            add_impl_getters!(must_get_agent_id, get_agent_id, ScAgentId);
            add_impl_getters!(must_get_address, get_address, ScAddress);
            add_impl_getters!(must_get_request_id, get_request_id, ScRequestId);
            add_impl_getters!(must_get_hname, get_hname, ScHname);
            add_impl_getters!(must_get_hash, get_hash, ScHash);
            add_impl_getters!(must_get_contract_id, get_contract_id, ScContractId);
            add_impl_getters!(must_get_color, get_color, ScColor);
            add_impl_getters!(must_get_chain_id, get_chain_id, ScChainId);
        }
    };
}

add_impl_getters!(ScFuncContext);
add_impl_getters!(ScViewContext);

pub fn require_if_needed<TContext:ScBaseContext>(context : &TContext, condition : bool, error_message : &str) {
    context.require(condition, error_message);
}