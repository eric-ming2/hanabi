fn main() {
    let out_dir = "src/generated";

    std::fs::create_dir_all(out_dir).unwrap();

    prost_build::Config::new()
        .out_dir(out_dir) // Set the output directory
        .compile_protos(
            &["../proto/requests.proto", "../proto/responses.proto"],
            &["../proto"],
        )
        .unwrap();
}
