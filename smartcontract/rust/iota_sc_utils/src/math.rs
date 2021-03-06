use wasmlib::*;

pub trait SafeMath {
    fn safe_add<TContext: ScBaseContext>(&self, b : &Self, ctx : &TContext) -> Self where Self : Sized;
    fn safe_sub<TContext: ScBaseContext>(&self, b : &Self, ctx : &TContext) -> Self where Self : Sized;
    fn safe_mul<TContext: ScBaseContext>(&self, b : &Self, ctx : &TContext) -> Self where Self : Sized;
    fn safe_div<TContext: ScBaseContext>(&self, b : &Self, ctx : &TContext) -> Self where Self : Sized;
}

macro_rules! add_impl {
    ($t:ty) => {
        impl SafeMath for $t {
            
            fn safe_add<TContext : ScBaseContext>(&self, b: &$t, ctx : &TContext) -> $t where Self : Sized {
                let a_plus_b = self.checked_add(*b);
                match a_plus_b {
                    Some(valid_result) => return valid_result,
                    _ => {
                        const INVALID_ADDITION : &str = "Invalid addition operation";
                        ctx.panic(INVALID_ADDITION);
                        panic!(INVALID_ADDITION)
                    },
                }
            }

            fn safe_sub<TContext : ScBaseContext>(&self, b: &$t, ctx : &TContext) -> $t where Self : Sized {
                let a_minus_b = self.checked_sub(*b);
                match a_minus_b {
                    Some(valid_result) => return valid_result,
                    _ => {
                        const INVALID_SUBTRACTION : &str = "Invalid subtraction operation";
                        ctx.panic(INVALID_SUBTRACTION);
                        panic!(INVALID_SUBTRACTION)
                    },
                }
            }

            fn safe_mul<TContext : ScBaseContext>(&self, b: &$t, ctx : &TContext) -> $t where Self : Sized {
                let a_times_b = self.checked_mul(*b);
                match a_times_b {
                    Some(valid_result) => return valid_result,
                    _ => {
                        const INVALID_MULTIPLICATION : &str = "Invalid multiplication operation";
                        ctx.panic(INVALID_MULTIPLICATION);
                        panic!(INVALID_MULTIPLICATION)
                    },
                }
            }

            fn safe_div<TContext : ScBaseContext>(&self, b: &$t, ctx : &TContext) -> $t where Self : Sized {
                let a_divided_by_b = self.checked_div(*b);
                match a_divided_by_b {
                    Some(valid_result) => return valid_result,
                    None => {
                        const INVALID_DIVISION : &str = "Invalid division operation. Reason: Division by zero.";
                        ctx.panic(INVALID_DIVISION);
                        panic!(INVALID_DIVISION)
                    }
                }
            }
        }
    };

    ($t1:ty, $t2:ty, $t3:ty, $t4:ty, $t5:ty) => {
        add_impl!($t1);
        add_impl!($t2);
        add_impl!($t3);
        add_impl!($t4);
        add_impl!($t5);
    };
}

add_impl!(u8, u16, u32, u64, usize);
add_impl!(i8, i16, i32, i64, isize);