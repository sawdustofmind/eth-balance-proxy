#/bin/bash
# This script sends 10,000 requests to the /eth/balance endpoint for each address in the list.

count=10

addresses=(
  # real addresses
  0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae
  0x742d35cc6634c0532925a3b844bc454e4438f44e
  0x53d284357ec70ce289d6d64134dfac8e511c8a3d
  0xfe9e8709d3215310075d67e3ed32a380ccf451c8
  0x66f820a414680b5bcda5eeca5dea238543f42054
  0xf977814e90da44bfa03b6295a0616a897441acec
  0x281055afc982d96fab65b3a49cac8b878184cb16
  0x6f46cf5569aefa1acc1009290c8e043747172d89
  0x5abfec25f74cd88437631a7731906932776356f9
  0x267be1c1d684f78cb4f6a176c4911b741e4ffdc0
  0x$(openssl rand -hex 20) # Random address for testing
)

for i in $(seq 1 $count); do
  for addr in "${addresses[@]}"; do
    curl "localhost:8080/eth/balance/$addr"
  done
done
