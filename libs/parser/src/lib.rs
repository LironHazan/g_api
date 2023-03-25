#![allow(clippy::type_complexity)]

// reusing my parser from: https://github.com/LironHazan/dummy-ql-example/tree/master

use nom::bytes::complete::{tag, escaped_transform};
use nom::multi::many0;
use nom::sequence::{delimited, tuple};
use nom::multi::separated_list0;
use nom::{IResult};
use nom::character::complete::none_of;
use nom::combinator::{recognize, map_parser};
use nom::error::{Error};
use nom::branch::alt;

type BaseOutput<'a> = (&'a str, &'a str, &'a str, &'a str, String);

fn parse_quoted(input: &str) -> IResult<&str, String> {
  let seq = recognize(separated_list0(tag("\"\""), many0(none_of("\""))));
  let unquote = escaped_transform(none_of("\""), '\"', tag("\""));
  let res = delimited(tag("\""), map_parser(seq, unquote), tag("\""))(input)?;

  Ok(res)
}


fn base_parser<'a>() -> impl FnMut(&'a str) -> IResult<&'a str, BaseOutput, Error<&'a str>> {
  let f = tag("foo");
  let eq_op = tag("==");
  let value = parse_quoted;
  tuple((f, tag(" "), eq_op, tag(" "), value))
}

// The main parser/entry point
pub fn dummy_parser(
  i: &str,
) -> IResult<&str, (BaseOutput, &str, &str, (&str, BaseOutput)), Error<&str>> {
  let and_parser = tag("&&");
  let or_parser = tag("||");
  let and_or_choice = alt((and_parser, or_parser));
  let sub_parser = tuple((tag(" "), base_parser()));
  tuple((base_parser(), tag(" "), and_or_choice, sub_parser))(i)
}



#[cfg(test)]
mod tests {
    use super::*;

  #[test]
  fn test_dummy_parser() {
    let (_, res) = dummy_parser("foo == \"jj\" && foo == \"bazz\"").unwrap();
    let exp_tuple = (
      (
        "foo",
        " ",
        "==",
        " ",
        "jj",
      ),
      " ",
      "&&",
      (
        " ",
        (
          "foo",
          " ",
          "==",
          " ",
          "bazz",
        ),
      ),
    );

    println!("res {:#?}: ",  res);

    assert_eq!(res.1, exp_tuple.1);
  }
}
