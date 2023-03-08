use parser::dummy_parser;

fn main() {
  match dummy_parser("foo == \"jj\" && foo == \"bazz\"") {
    Ok((remaining_input, output)) => {
      // Handle successful parsing result
      let (base_output, str1, str2, (str3, base_output2)) = output;
      println!("Parsing succeeded");
      println!("Remaining input: {:?}", remaining_input);
      println!("Base output: {:?}", base_output);
      println!("Str1: {:?}", str1);
      println!("Str2: {:?}", str2);
      println!("Str3: {:?}", str3);
      println!("Base output 2: {:?}", base_output2);
    }
    Err(_) => {
      println!("Parsing incomplete");
    }
  }
}
