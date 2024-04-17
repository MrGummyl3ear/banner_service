import http from 'k6/http';

export const options = {
    vus: 1,
    iterations: 1,
};

export default function () {
    const url = 'http://localhost:8000/banner';
    const authToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTMzOTEwMTMsImlhdCI6MTcxMzM0NzgxMywicm9sZXMiOlsiQWRtaW4iXX0.gRHtrFxguw3Hg6KenLrGRLq_X7w9gBnwCYILu_JzTgk';
    const headers = {
        Authorization: `Bearer ${authToken}`
    };

    const response = http.get(url, { headers: headers });

    if (response.status === 200) {
        console.log(response.body);
    } else {
        console.error(`Error getting data banners: ${response.status}`);
    }
}