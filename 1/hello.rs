use std::f64::consts;

fn double_num(num: f64) -> f64{
   return num * 2.0;
}

fn by_ref(num: &i32) -> i32{
    return *num * 3;
}
fn by_mut_ref(num: &mut i32){
    *num *= 2;
}

fn tuloop(){
    let max = 10;
    let mut age = 1;
    for i in 0..max{

        let even_odd = if i % 2 == 0 {"even"} else {"odd"};
        age += 1;
        println!("[{}] Hello kayi {}",even_odd, i);
    }

    println!("\nAt the end, she was {} years old", age);
    println!("But {} years old when doubled", double_num(age as f64));
    println!("{}", age); // Did not change
}

fn new_line(){println!("\n");}
fn main(){

    let mut age = 30;
    println!("Hello fools! Age: {}", age);
    tuloop();
    let tripple_age = by_ref(&age);
    println!("\nTrippled age: {}", tripple_age);
    println!("Old age: {}", age);

    by_mut_ref(&mut age);
    println!("\nDouble the age");
    println!("Old age: {}", age);
    new_line();
    println!("PI = {}", consts::PI);
}