use wasmlib::*;
use crate::getter::Getter;
use crate::getter::PARAMS;

// Primitive types
add_impl_pub_fns!(PARAMS, must_get_string, get_string, String);
add_impl_pub_fns!(PARAMS, must_get_int64, get_int64, i64);
add_impl_pub_fns!(PARAMS, must_get_bytes, get_bytes, Vec<u8>);

// ISCP Types
add_impl_pub_fns!(PARAMS, must_get_agent_id, get_agent_id, ScAgentId);
add_impl_pub_fns!(PARAMS, must_get_address, get_address, ScAddress);
add_impl_pub_fns!(PARAMS, must_get_request_id, get_request_id, ScRequestId);
add_impl_pub_fns!(PARAMS, must_get_hname, get_hname, ScHname);
add_impl_pub_fns!(PARAMS, must_get_hash, get_hash, ScHash);
add_impl_pub_fns!(PARAMS, must_get_contract_id, get_contract_id, ScContractId);
add_impl_pub_fns!(PARAMS, must_get_color, get_color, ScColor);
add_impl_pub_fns!(PARAMS, must_get_chain_id, get_chain_id, ScChainId);