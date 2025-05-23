// Configuration for different environments
const config = {
    // Local development
    development: {
        API_BASE_URL: 'http://localhost:8080'
    },
    // GitHub Pages with external API
    production: {
        API_BASE_URL: 'https://your-api.railway.app'  // Replace with your actual API URL
    }
};

// Detect environment
const isDevelopment = window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1';
const currentConfig = isDevelopment ? config.development : config.production;

// Export for use in other scripts
window.APP_CONFIG = currentConfig; 