use wasmlib::*;

macro_rules! add_impl_pub_fns {
    ($must_get_func_name:ident, $t1:ty) => {
        /// Tries to get a variable. Panics if it can't find it.
        pub fn $must_get_func_name<TGetVar:GetVar>(variable_name : &str, ctx: &TGetVar) -> $t1 {
            ctx.$must_get_func_name(variable_name)
        }
    };

    ($get_func_name:ident, $t1:ty) => {
        /// Tries to get a variable. Returns default value if it can't find it.
        pub fn $get_func_name<TGetVar:GetVar>(variable_name : &str, ctx: &TGetVar) -> $t1 {
            ctx.$get_func_name(variable_name)
        }
    };

    ($must_get_func_name:ident, $get_func_name:ident, $t1:ty) => {
        add_impl_pub_fns!($must_get_func_name, $t1);
        add_impl_pub_fns!($get_func_name, $t1);
    };
}

// Basic types
add_impl_pub_fns!(must_get_string, get_string, String);
add_impl_pub_fns!(must_get_i64, get_i64, i64);
add_impl_pub_fns!(must_get_bytes, get_bytes, Vec<u8>);
// ISCP Types
add_impl_pub_fns!(must_get_agent_id, get_agent_id, ScAgentId);
add_impl_pub_fns!(must_get_address, get_address, ScAddress);
add_impl_pub_fns!(must_get_request_id, get_request_id, ScRequestId);
add_impl_pub_fns!(must_get_hname, get_hname, ScHname);
add_impl_pub_fns!(must_get_hash, get_hash, ScHash);
add_impl_pub_fns!(must_get_contract_id, get_contract_id, ScContractId);
add_impl_pub_fns!(must_get_color, get_color, ScColor);
add_impl_pub_fns!(must_get_chain_id, get_chain_id, ScChainId);

macro_rules! add_impl_get_var {
    ($t1:ty, $t2:ty) => {
        impl GetVar for $t1 {
            
            fn must_get_string(&self, variable_name : &str) -> String {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_string(variable_name);
                self.require_if_needed(variable.exists(), "string variable not found");
                variable.value()
            }

            fn get_string(&self, variable_name : &str) -> String {
                let getter : $t2 = self.load_getter();
                getter.get_string(variable_name).value()
            }

            fn must_get_i64(&self, variable_name : &str) -> i64 {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_int64(variable_name);
                self.require_if_needed(variable.exists(), "i64 variable not found");
                variable.value()
            }

            fn get_i64(&self, variable_name : &str) -> i64 {
                let getter : $t2 = self.load_getter();
                getter.get_int64(variable_name).value()
            }

            fn must_get_bytes(&self, variable_name : &str) -> Vec<u8> {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_bytes(variable_name);
                self.require_if_needed(variable.exists(), "bytes parameter not found");
                variable.value()
            }

            fn get_bytes(&self, variable_name : &str) -> Vec<u8> {
                let getter : $t2 = self.load_getter();
                getter.get_bytes(variable_name).value()
            }

            fn must_get_agent_id(&self, variable_name : &str) -> ScAgentId {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_agent_id(variable_name);
                self.require_if_needed(variable.exists(), "agent_id variable not found");
                variable.value()
            }

            fn get_agent_id(&self, variable_name : &str) -> ScAgentId {
                let getter : $t2 = self.load_getter();
                getter.get_agent_id(variable_name).value()
            }

            fn must_get_address(&self, variable_name : &str) -> ScAddress {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_address(variable_name);
                self.require_if_needed(variable.exists(), "address variable not found");
                variable.value()
            }
            
            fn get_address(&self, variable_name : &str) -> ScAddress {
                let getter : $t2 = self.load_getter();
                getter.get_address(variable_name).value()
            }

            fn must_get_request_id(&self, variable_name : &str) -> ScRequestId {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_request_id(variable_name);
                self.require_if_needed(variable.exists(), "request_id variable not found");
                variable.value()
            }
            
            fn get_request_id(&self, variable_name : &str) -> ScRequestId {
                let getter : $t2 = self.load_getter();
                getter.get_request_id(variable_name).value()
            }

            fn must_get_hname(&self, variable_name : &str) -> ScHname {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_hname(variable_name);
                self.require_if_needed(variable.exists(), "hname variable not found");
                variable.value()
            }
            
            fn get_hname(&self, variable_name : &str) -> ScHname {
                let getter : $t2 = self.load_getter();
                getter.get_hname(variable_name).value()
            }

            fn must_get_hash(&self, variable_name : &str) -> ScHash {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_hash(variable_name);
                self.require_if_needed(variable.exists(), "hash variable not found");
                variable.value()
            }
            
            fn get_hash(&self, variable_name : &str) -> ScHash {
                let getter : $t2 = self.load_getter();
                getter.get_hash(variable_name).value()
            }

            fn must_get_contract_id(&self, variable_name : &str) -> ScContractId {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_contract_id(variable_name);
                self.require_if_needed(variable.exists(), "contract_id variable not found");
                variable.value()
            }
            
            fn get_contract_id(&self, variable_name : &str) -> ScContractId {
                let getter : $t2 = self.load_getter();
                getter.get_contract_id(variable_name).value()
            }

            fn must_get_color(&self, variable_name : &str) -> ScColor {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_color(variable_name);
                self.require_if_needed(variable.exists(), "color variable not found");
                variable.value()
            }
            
            fn get_color(&self, variable_name : &str) -> ScColor {
                let getter : $t2 = self.load_getter();
                getter.get_color(variable_name).value()
            }

            fn must_get_chain_id(&self, variable_name : &str) -> ScChainId {
                let getter : $t2 = self.load_getter();
                let variable = getter.get_chain_id(variable_name);
                self.require_if_needed(variable.exists(), "chain_id variable not found");
                variable.value()
            }
            
            fn get_chain_id(&self, variable_name : &str) -> ScChainId {
                let getter : $t2 = self.load_getter();
                getter.get_chain_id(variable_name).value()
            }
        }

        impl Getter<$t2> for $t1 {
            fn load_getter(&self) -> $t2 {
                self.state()
            }
        
            fn require_if_needed(&self, condition : bool, error_message : &str){
                self.require(condition, error_message)
            }
        }
    };
}

macro_rules! add_GetVar_fns {
    ($must_get_func_name:ident, $get_func_name:ident, $t1:ty) => {
        /// Tries to get a variable. Panics if it can't find it.
        fn $must_get_func_name(&self, variable_name : &str) -> $t1;
        /// Tries to get a variable. Returns default value if it can't find it.
        fn $get_func_name(&self, variable_name : &str) -> $t1;
    };

    ($get_func_name : ident, $map:ty) => {
        fn $get_func_name(&self, variable_name : &str) -> $map;
    };
}

pub trait GetVar {
    // Primitive types
    add_GetVar_fns!(must_get_string, get_string, String);
    add_GetVar_fns!(must_get_i64, get_i64, i64);
    add_GetVar_fns!(must_get_bytes, get_bytes, Vec<u8>);

    // ISCP types
    add_GetVar_fns!(must_get_agent_id, get_agent_id, ScAgentId);
    add_GetVar_fns!(must_get_address, get_address, ScAddress);
    add_GetVar_fns!(must_get_request_id, get_request_id, ScRequestId);
    add_GetVar_fns!(must_get_hname, get_hname, ScHname);
    add_GetVar_fns!(must_get_hash, get_hash, ScHash);
    add_GetVar_fns!(must_get_contract_id, get_contract_id, ScContractId);
    add_GetVar_fns!(must_get_color, get_color, ScColor);
    add_GetVar_fns!(must_get_chain_id, get_chain_id, ScChainId);
}

trait Getter<TMap> {
    fn load_getter(&self) -> TMap;
    fn require_if_needed(&self, condition : bool, error_message : &str);
}

add_impl_get_var!(ScFuncContext, ScMutableMap);
add_impl_get_var!(ScViewContext, ScImmutableMap);

// ----

macro_rules! add_impl_get_map {
    ($context:ty, $get_func_name:ident, $map:ty) => {
        impl GetMap<$map> for $context {
            fn $get_func_name(&self, variable_name : &str) -> $map {
                let getter = self.load_getter();
                getter.get_map(variable_name)
            }
        }
    };
}

trait GetMap<TMap> {
    fn get_map(&self, variable_name : &str) -> TMap;
}

add_impl_get_map!(ScFuncContext, get_map, ScMutableMap);
add_impl_get_map!(ScViewContext, get_map, ScImmutableMap);