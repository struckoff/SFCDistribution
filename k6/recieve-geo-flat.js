import http from "k6/http";
import { sleep } from "k6";
import { group } from "k6";
import { check } from "k6";


let iters = 10
let vus = 1




let p_host = "http://localhost:42080/extract/alldata-0"
let dataFile = open(__ENV["K6_KEYS"]);

// let keySeqLen = 100


export let options = {
    // tags: {
    //     "name": "receive geo"
    // },
    // setupTimeout: "30m",
    // maxDuration: "10s",
    scenarios: {
        recieve_1m: {
          executor: 'per-vu-iterations',  // name of the executor to use
          // common scenario configuration
        //   startTime: '10s',
          gracefulStop: '5s',
          env: { INTERVAL: '1m' },
          vus: vus,
          iterations: iters,
          maxDuration: '10h',
        },
      }
//   minIterationDuration: "100ms"
};
export function setup() {
    var url = 
    return url
}

export default function(url) {
    var res = http.get(url, null, {tags: {name: 'get_download_geo'}});
    if (res.status >= 400){
        console.error(res.body)
    }
    check(res, {
        "is status OK": (r) => r.status < 400,
        "is status not 404": (r) => r.status != 404,
        "is status not 403": (r) => r.status != 403,
        "is status not 500": (r) => r.status != 500,
        "is status not 503": (r) => r.status != 503,
    });
}
