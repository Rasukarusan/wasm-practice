function validProof(lastProof, proof, prefix) {
  const guess = `${lastProof}${proof}`
  const shaObj = new jsSHA("SHA-256", "TEXT", { encoding: "UTF8" });
  shaObj.update(guess);
  const hash = shaObj.getHash("HEX");
  return hash.startsWith(prefix)
}

function proofOfWork(lastProof, prefix) {
  let p = 0
  while(!validProof(lastProof, p, prefix)) {
    p++
  }
  return p
}

function mineByJs(value) {
  const prefix = "0".repeat(value)
  const startTime = performance.now(); // 開始時間
  const result = proofOfWork(100, prefix)
  const endTime = performance.now(); // 終了時間
  document.getElementById("time").innerHTML = (endTime - startTime).toPrecision(6) + "ms"
  return result
}

