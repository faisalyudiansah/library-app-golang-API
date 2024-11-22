import http from 'k6/http';
import { sleep } from 'k6';

export const options = {
  stages: [
    { duration: '5m', target: 200 }, // ramp-up
    { duration: '20m', target: 200 }, // stable
    { duration: '5m', target: 0 }, // ramp-down
  ],
  thresholds: {
    http_req_duration: ['p(99)<100'], // 99% of requests must complete within 100ms
  }
}

export default function () {
  http.get('http://localhost:8081/books');
  sleep(1);
}
