const express = require('express');
const client = require('prom-client');
const app = express();
const PORT = process.env.PORT || 3000;

//Prometheus metrics

const register = new client.Registry();
client.collectDefaultMetrics({ register });

const httpRequestCounter = new client.Counter({
	name: 'http_request_total',
	help: 'Total number of HTTP requests',
	labelNames: ['method', 'route', 'status_code'],
	register: [register]
});

//Middleware
app.use(express.json());

//Health check
app.get('/health', (req, res) => {
	httpRequestCounter.inc({method: 'GET', route: '/health', status_code: 200});
	res.json({ status: 'healthy', service: 'node-api', timestamp: new Date().toISOString() });
});

//Main endpoint
app.get('/api/hello', (req, res) => {
	httpRequestCounter.inc({ method: 'GET', route: '/api/hello', status_code: 200 });
	res.json({
		message: 'Hello from Node.js API 🫡',
		version: '1.0.0',
		language: 'Node.js'});
});

//Metrics endpoints for Prometheus
app.get('/metrics', async (req, res) => {
	res.set('Content-Type', register.contentType);
	res.end(await register.metrics());
});

app.listen(PORT, () => {
	console.log( `Node.js API running on port ${PORT}`);
});
module.exports = app;


