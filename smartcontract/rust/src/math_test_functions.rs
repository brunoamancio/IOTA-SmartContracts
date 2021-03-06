use wasmlib::{ScBaseContext, ScExports};
use iota_sc_utils::math::SafeMath;

pub trait SafeMathScFunction {
    fn safe_add_no_overflow_function<TContext: ScBaseContext>(ctx : &TContext);
    fn safe_add_with_overflow_function<TContext: ScBaseContext>(ctx : &TContext);

    fn safe_sub_no_overflow_function<TContext: ScBaseContext>(ctx : &TContext);
    fn safe_sub_with_overflow_function<TContext: ScBaseContext>(ctx : &TContext);

    fn safe_mul_no_overflow_function<TContext: ScBaseContext>(ctx : &TContext);
    fn safe_mul_with_overflow_function<TContext: ScBaseContext>(ctx : &TContext);

    fn safe_div_no_overflow_function<TContext: ScBaseContext>(ctx : &TContext);
    fn safe_div_with_overflow_function<TContext: ScBaseContext>(ctx : &TContext);
}

macro_rules! add_impl {
    ($trait_name:ident, $t:ident) => {
        impl $trait_name for $t {
            
            // -------------------------------------  ADDITION -------------------------------------

            fn safe_add_no_overflow_function<TContext : ScBaseContext>(ctx : &TContext) {
                let value_to_calc : $t = 0;
                let max : $t = std::$t::MAX;
                // Do not cause an overflow
                let _result = max.safe_add(&value_to_calc, ctx);

                ctx.log(&format!("Function \"{}\" finished without an overflow)", stringify!(safe_add_no_overflow_function)));
            }

            fn safe_add_with_overflow_function<TContext : ScBaseContext>(ctx : &TContext) {
                let value_to_calc : $t = 1;
                let max : $t = std::$t::MAX;
                // Cause an overflow
                let _result = max.safe_add(&value_to_calc, ctx);

                ctx.log(&format!("Function \"{}\" finished with an overflow)", stringify!(safe_add_with_overflow_function)));
            }

            // -------------------------------------  SUBTRACTION -------------------------------------

            fn safe_sub_no_overflow_function<TContext : ScBaseContext>(ctx : &TContext) {
                let value_to_calc : $t = 0;
                let min : $t = std::$t::MIN;
                // Do not cause an overflow
                let _result = min.safe_sub(&value_to_calc, ctx);

                ctx.log(&format!("Function \"{}\" finished without an overflow)", stringify!(safe_sub_no_overflow_function)));
            }

            fn safe_sub_with_overflow_function<TContext : ScBaseContext>(ctx : &TContext) {
                let value_to_calc : $t = 1;
                let min : $t = std::$t::MIN;
                // Cause an overflow
                let _result = min.safe_sub(&value_to_calc, ctx);

                ctx.log(&format!("Function \"{}\" finished with an overflow)", stringify!(safe_sub_with_overflow_function)));
            }

            // -------------------------------------  MULTIPLICATION -------------------------------------

            fn safe_mul_no_overflow_function<TContext : ScBaseContext>(ctx : &TContext) {
                let value_to_calc : $t = 1;
                let max : $t = std::$t::MAX;
                // Do not cause an overflow
                let _result = max.safe_mul(&value_to_calc, ctx);

                ctx.log(&format!("Function \"{}\" finished without an overflow)", stringify!(safe_mul_no_overflow_function)));
            }

            fn safe_mul_with_overflow_function<TContext : ScBaseContext>(ctx : &TContext) {
                let value_to_calc : $t = 2;
                let max : $t = std::$t::MAX;
                // Cause an overflow
                let _result = max.safe_mul(&value_to_calc, ctx);

                ctx.log(&format!("Function \"{}\" finished with an overflow)", stringify!(safe_mul_with_overflow_function)));
            }

            // -------------------------------------  DIVISION -------------------------------------

            fn safe_div_no_overflow_function<TContext : ScBaseContext>(ctx : &TContext) {
                let value_to_calc : $t = 1;
                let max : $t = std::$t::MAX;
                // Do not cause an overflow
                let _result = max.safe_div(&value_to_calc, ctx);

                ctx.log(&format!("Function \"{}\" finished without an overflow)", stringify!(safe_div_no_overflow_function)));
            }

            fn safe_div_with_overflow_function<TContext : ScBaseContext>(ctx : &TContext) {
                let value_to_calc : $t = 0;
                let min : $t = std::$t::MIN;
                // Cause an overflow
                let _result = min.safe_div(&value_to_calc, ctx);

                ctx.log(&format!("Function \"{}\" finished with an overflow)", stringify!(safe_div_with_overflow_function)));
            }
        }
    };

    ($trait_name:ident, $t1:ident, $t2:ident, $t3:ident, $t4:ident, $t5:ident) => {
        add_impl!($trait_name, $t1);
        add_impl!($trait_name, $t2);
        add_impl!($trait_name, $t3);
        add_impl!($trait_name, $t4);
        add_impl!($trait_name, $t5);
    };
}

add_impl!(SafeMathScFunction, u8, u16, u32, u64, usize);
add_impl!(SafeMathScFunction, i8, i16, i32, i64, isize);

macro_rules! register_func {
    ($sc_exports:tt, $t:ident) => {
        $sc_exports.add_func(&format!("{}_{}", stringify!($t), stringify!(safe_add_no_overflow_function)), $t::safe_add_no_overflow_function);
        $sc_exports.add_func(&format!("{}_{}", stringify!($t), stringify!(safe_add_with_overflow_function)), $t::safe_add_with_overflow_function);

        $sc_exports.add_func(&format!("{}_{}", stringify!($t), stringify!(safe_sub_no_overflow_function)), $t::safe_sub_no_overflow_function);
        $sc_exports.add_func(&format!("{}_{}", stringify!($t), stringify!(safe_sub_with_overflow_function)), $t::safe_sub_with_overflow_function);

        $sc_exports.add_func(&format!("{}_{}", stringify!($t), stringify!(safe_mul_no_overflow_function)), $t::safe_mul_no_overflow_function);
        $sc_exports.add_func(&format!("{}_{}", stringify!($t), stringify!(safe_mul_with_overflow_function)), $t::safe_mul_with_overflow_function);

        $sc_exports.add_func(&format!("{}_{}", stringify!($t), stringify!(safe_div_no_overflow_function)), $t::safe_div_no_overflow_function);
        $sc_exports.add_func(&format!("{}_{}", stringify!($t), stringify!(safe_div_with_overflow_function)), $t::safe_div_with_overflow_function);
    };

    ($sc_exports:tt, $t1:ident, $t2:ident, $t3:ident, $t4:ident, $t5:ident) => {
        register_func!($sc_exports, $t1);
        register_func!($sc_exports, $t2);
        register_func!($sc_exports, $t3);
        register_func!($sc_exports, $t4);
        register_func!($sc_exports, $t5);
    };
}

// Every registration is a call to sc_exports.add_func for all safe math functions for the specified type. 
// This allows testing the SCFunctions from the unit tests on Solo
pub fn register_math_sc_functions(sc_exports : &ScExports){
    register_func!(sc_exports, u8, u16, u32, u64, usize);
    register_func!(sc_exports, i8, i16, i32, i64, isize);
}