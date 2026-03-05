// Copyright (C) 2022 Explore.dev, Unipessoal Lda - All Rights Reserved
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

fn main() {
    tonic_prost_build::configure()
        .compile_protos(&["pb/services/semantic.proto"],
                        &["pb"]).unwrap();
}