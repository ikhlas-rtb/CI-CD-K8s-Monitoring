from flask import Flask, jsonify
from prometheus_client import Counter, generate_latest, REGISTRY
import time

app = Flask(__name__)

http_requests_total = Counter(
	'http_requests_total',
	'Total HTTP requests',
	['method', 'route', 'status_code']
)

@app.route('/health', methods=['GET'])
def health():
	http_requests_total.labels(method='GET', route='/health', status_code=200).inc()
	return jsonify({
		'status': 'healthy',
		'service': 'python-api',
		'timestamp': time.time()
}), 200
@app.route('/api/hello', methods=['GET'])
def hello():
	http_requests_total.labels(method='GET', route='/api/hello', status_code=200).inc()
	return jsonify({
		'message': 'Hello from Flask API! 🫡',
		'version': '1.0.0',
		'language': 'python/Flask'

}), 200
@app.route('/metrics', methods=['GET'])
def metrics():
	return generate_latest(REGISTRY), 200, {'Content-Type': 'text/plain; charset=utf-8'}
if __name__ == '__main__':
	app.run(host='0.0.0.0', port=5000, debug=False)
