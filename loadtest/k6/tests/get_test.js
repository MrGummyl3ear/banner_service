import { sleep } from 'k6';
import http from 'k6/http';
import { randomIntBetween, randomItem } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

const banners = JSON.parse(open(`./banners.json`));


export const options = {
    scenarios: {
        load_test: {
            executor: 'constant-arrival-rate',
            rate: 300,
            timeUnit: '1s',
            duration: '1m',
            preAllocatedVUs: 10,
            maxVUs: 40,
        },
    },
    thresholds: {
        http_req_failed: ['rate<0.0001'],
        http_req_duration: ['p(95)<200'],
    },
};

export default function () {

    const authToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTM0ODI3OTgsImlhdCI6MTcxMzQzOTU5OCwicm9sZXMiOlsiQWRtaW4iXX0.GQZdQNgcns574jzQDaQj5ATEI_I5DM129FAZxQ7Ro0Y';

    const banner = randomItem(banners);
    const tagID = randomItem(banner.tag_ids);
    const limit = randomIntBetween(0, 30);
    const offset = randomIntBetween(0, 30);
    const pattern = randomIntBetween(1,3);

    let url = http.url``;
    if (pattern == 1){
        url = `http://localhost:8000/banner?feature_id=${banner.feature_id}&tag_id=${tagID}&limit=${limit}&offset=${offset}`;
    }
    if (pattern == 2){
        url = `http://localhost:8000/banner?feature_id=${banner.feature_id}&limit=${limit}&offset=${offset}`;
    }   
    if (pattern == 3){
        url = `http://localhost:8000/banner?limit=${limit}&offset=${offset}`;
    }
    const params = {
        headers: {
            Authorization: `Bearer ${authToken}`
        },
    };

    http.get(url, params);

}