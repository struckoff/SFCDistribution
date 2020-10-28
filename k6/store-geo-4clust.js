import http from "k6/http";
import { sleep } from "k6";
import { group } from "k6";
import { check } from "k6";


let iters = __ENV["K6_ITERATIONS"] * 1
let vus = __ENV["K6_VUS"] * 1

// let p_host = "http://ip-172-31-42-150.eu-central-1.compute.internal:47375"
let p_hosts = [
    "http://localhost:9190",
    // "http://localhost:9191",
    // "http://localhost:9192",
    // "http://localhost:9193",
    // "http://localhost:9194",
    // "http://localhost:9195",
    // "http://localhost:9196",
    // "http://localhost:9197",
    // "http://localhost:9198",
    // "http://localhost:9199",
    // "http://localhost:9200",
]
// let dataFile = open("/home/struckoff/Documents/alldata/cut_10k.csv");

let minLat = -90
let maxLat = 90
let minLon = -180
let maxLon = 180
let gap = 30

let cMinLat = minLat * .75
let cMaxLat = maxLat * .75
let cMinLon = minLon * .75
let cMaxLon = maxLon * .75

let step = .5

export let options = {
    scenarios: {
        recieve_geo_scenario: {  // arbitrary and unique scenario name, will appear in the result summary, tags, etc.
          executor: 'per-vu-iterations',  // name of the executor to use
          // common scenario configuration
        //   startTime: '10s',
          gracefulStop: '5s',
          // executor-specific configuration
          vus: vus,
          iterations: iters,
          maxDuration: '10h',
        },
      }
};

function randomPoint(){
    // var timeInMs = Date.now() - Math.floor(Math.random());
    var timeInMs = Date.now();
    var val = Math.random() + 10
    return timeInMs + ';' + val
}

function genData(){
    var world = {LatMin:minLat, LatMax:maxLat, LonMin:minLon, LonMax:maxLon, Color:"ffffff", Name:"world"} // all geo

    var clusters = [
        {LatMin:gap, LatMax:cMaxLat, LonMin:cMinLon, LonMax:-gap, Color:"508c36", Name:`${gap}<=lat<=${maxLat} ${minLon}<=lon<=-${gap}`}, // 1 cluster
        {LatMin:gap, LatMax:cMaxLat, LonMin:gap, LonMax:cMaxLon, Color:"2d728f", Name:`${gap}<=lat<=${maxLat} ${gap}<=lon<=${maxLon}`}, // 2 cluster
        {LatMin:cMinLat, LatMax:-gap, LonMin:cMinLon, LonMax:-gap, Color:"f7d08a", Name:`${minLat}<=lat<=-${gap} ${minLon}<=lon<=-${gap}`}, // 3 cluster
        {LatMin:cMinLat, LatMax:-gap, LonMin:gap, LonMax:cMaxLon, Color:"ed6da0", Name:`${minLat}<=lat<=-${gap} ${gap}<=lon<=${maxLon}`}, // 4 cluster
    ]

    var iter = 0;
    var data = [];
    for (var lat = minLat; lat <maxLat; lat+= step) {
        for (var lon = minLon; lon <maxLon; lon+= step){
            var rate = Math.random()
            if (rate < 1/3){
                data[iter] = genPoint(world)
            } else {
                var params = clusters[Math.floor(Math.random() * clusters.length)]
                data[iter] = genPoint(params)
            }
            iter++
        }
    }
    return data
}

function genPoint(params){
    var lon = Math.random() * (params.LonMax - params.LonMin) + params.LonMin
    var lat = Math.random() * (params.LatMax - params.LatMin) + params.LatMin
    if (minLon > lon && lon < maxLon){
        console.error("ERRRR lon", lon)
    }
    if (minLat > lat && lat < maxLat){
        console.error("ERRRR lat", lat)
    }
    // return JSON.stringify({"Lon":lon, "Lat":lat})
    
    return JSON.stringify({"Lon":lon, "Lat":lat, "Cluster":params.Name, "ClusterColor":params.Color})
}

export function setup() {
    var points = genData()
    // var index = 0
    // while (index < 10000){
    //     yield points[index]
    //     index++
    // }
    return points
}

export default function(points) {
    let idx = ((iters * (__VU - 1)) + __ITER) % (points.length)
    let key = points[idx]

    // console.log(key, idx, __ITER, __VU, maxIter, points.length)
    var res = http.post(p_hosts[__VU % p_hosts.length] + "/put/"+key, null, {tags: {name: 'post_upload_geo'}});
    if (res.status >= 400){
        console.error(p_hosts[__VU % p_hosts.length], res.body, key)
    }
    check(res, {
        "is status OK": (r) => r.status < 400,
        "is status not 404": (r) => r.status != 404,
        "is status not 403": (r) => r.status != 403,
        "is status not 500": (r) => r.status != 500,
        "is status not 503": (r) => r.status != 503,
    });
}