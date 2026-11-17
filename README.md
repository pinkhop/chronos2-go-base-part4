# chronos2-go-base-part4

Shards 10 through 12 of the 12 that make up Amazon's [Chronos-2](https://huggingface.co/amazon/chronos-2) checkpoint, for [`chronos2-go-base`](https://github.com/pinkhop/chronos2-go-base).

Do not import this module. Import `github.com/pinkhop/chronos2-go-base`, which requires all four parts and assembles them into a ready-to-use model; `go get` fetches the parts automatically. The checkpoint is split across four modules so that each stays small enough for the Go module proxy to serve.


## Contents

| File | Bytes | SHA-256 |
|---|---|---|
| `model-00010-of-00012.safetensors` | 37,759,488 | `970cc12d662c1cd39dd774d494f5dc6494009f7d148784a666dc1228b48c7bf4` |
| `model-00011-of-00012.safetensors` | 37,759,488 | `f9028524d4fa1b25b0b765fa69f4fb20fd1c324ff26aa541bdc09086fb55a4ff` |
| `model-00012-of-00012.safetensors` | 43,691,376 | `ba95d3479944e7126ff2415631caf4d9c13eb9cafa4f912acefc4acbc5412e5d` |

The shards are fp32 safetensors files cut from the upstream `model.safetensors` at revision [`29ec3766d36d6f73f0696f85560a422f50e8498c`](https://huggingface.co/amazon/chronos-2/tree/29ec3766d36d6f73f0696f85560a422f50e8498c), with every tensor bit-identical to upstream. [`weights/manifest.json`](weights/manifest.json) records the upstream source and each shard's size and SHA-256, and `make verify-weights` checks the files against it.


## Contributing

[`AGENTS.md`](AGENTS.md) describes the repository, the `make` targets, and the rules every change follows, for human and AI contributors alike.


## License

Licensed under the Apache License, Version 2.0; see [`LICENSE`](LICENSE).

Chronos-2 was created by Amazon, and its weights are distributed under Apache-2.0; see [`NOTICE`](NOTICE).
