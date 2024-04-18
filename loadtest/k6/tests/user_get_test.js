import { sleep } from 'k6';
import http from 'k6/http';
import { randomIntBetween, randomItem } from 'https://jslib.k6.io/k6-utils/1.2.0/index.js';

const banners = JSON.parse(open(`./banners.json`));


export const options = {
    scenarios: {
        load_test: {
            executor: 'constant-arrival-rate',
            duration: '5m',
            preAllocatedVUs: 10,

            rate: 1000,
            timeUnit: '1s',
            maxVUs: 40,
        },
    },
    discardResponseBodies: true,
    thresholds: {
        http_req_failed: ['rate<0.0001'],
        http_req_duration: ['p(95)<200'],
    },
};

export default function () {

    const authToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTM0ODI3OTgsImlhdCI6MTcxMzQzOTU5OCwicm9sZXMiOlsiQWRtaW4iXX0.GQZdQNgcns574jzQDaQj5ATEI_I5DM129FAZxQ7Ro0Y';

    const banner = randomItem(banners);

    const tagID = randomItem(banner.tag_ids);

    let use_last_revision = false;
    if (randomIntBetween(0, 100) < 10) {
        use_last_revision = true;
    }

    const url = `http://localhost:8000/user_banner?feature_id=${banner.feature_id}&tag_id=${tagID}&use_last_revision=${use_last_revision}`;
    const params = {
        headers: {
            Authorization: `Bearer ${authToken}`,
        },
    };

    http.get(url, params);

}