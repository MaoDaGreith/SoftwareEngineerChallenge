class PackCalculator {
    constructor() {
        this.initializeDefaultPacks();
        this.bindEvents();
    }

    initializeDefaultPacks() {
        const defaultPacks = [250, 500, 1000, 2000, 5000];
        this.setPackSizes(defaultPacks);
    }

    bindEvents() {
        document.getElementById('add-pack').addEventListener('click', () => {
            this.addPackSizeInput();
        });

        document.getElementById('use-defaults').addEventListener('click', () => {
            this.initializeDefaultPacks();
        });

        document.getElementById('calculate').addEventListener('click', () => {
            this.calculatePacks();
        });

        // Enter key shortcut
        document.getElementById('order-amount').addEventListener('keypress', (e) => {
            if (e.key === 'Enter') {
                this.calculatePacks();
            }
        });

        // Handle remove buttons
        document.addEventListener('click', (e) => {
            if (e.target.classList.contains('remove-pack')) {
                this.removePackSizeInput(e.target.closest('.pack-size-input'));
            }
        });
    }

    setPackSizes(packSizes) {
        const container = document.querySelector('.pack-sizes');
        container.innerHTML = '';
        
        packSizes.forEach(size => {
            this.addPackSizeInput(size);
        });
    }

    addPackSizeInput(value = '') {
        const container = document.querySelector('.pack-sizes');
        const div = document.createElement('div');
        div.className = 'pack-size-input';
        div.innerHTML = `
            <input type="number" placeholder="Pack size" min="1" value="${value}">
            <button type="button" class="remove-pack">×</button>
        `;
        container.appendChild(div);
    }

    removePackSizeInput(element) {
        const container = document.querySelector('.pack-sizes');
        if (container.children.length > 1) {
            element.remove();
        }
    }

    getPackSizes() {
        const inputs = document.querySelectorAll('.pack-size-input input');
        const packSizes = [];
        
        inputs.forEach(input => {
            const value = parseInt(input.value);
            if (value && value > 0) {
                packSizes.push(value);
            }
        });
        
        return packSizes;
    }

    async calculatePacks() {
        const orderAmount = parseInt(document.getElementById('order-amount').value);
        const packSizes = this.getPackSizes();

        this.hideAllMessages();

        if (!orderAmount || orderAmount <= 0) {
            this.showError('Please enter a valid order amount');
            return;
        }

        if (packSizes.length === 0) {
            this.showError('Please specify at least one pack size');
            return;
        }

        this.showLoading();

        try {
            const response = await fetch('/api/packs/calculate', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    pack_sizes: packSizes,
                    order_amount: orderAmount
                })
            });

            const data = await response.json();

            if (response.ok) {
                this.showResults(data.packs, orderAmount);
            } else {
                this.showError(data.error || 'An error occurred');
            }
        } catch (error) {
            this.showError('Failed to connect to the server');
        } finally {
            this.hideLoading();
        }
    }

    showResults(packs, orderAmount) {
        const breakdown = document.getElementById('pack-breakdown');
        const summary = document.getElementById('summary');
        const results = document.getElementById('results');

        breakdown.innerHTML = '';

        let totalItems = 0;
        let totalPacks = 0;

        // Sort by size (largest first)
        const sortedPacks = Object.entries(packs).sort((a, b) => parseInt(b[0]) - parseInt(a[0]));

        sortedPacks.forEach(([size, quantity]) => {
            const packDiv = document.createElement('div');
            packDiv.className = 'pack-item';
            packDiv.innerHTML = `
                <div class="pack-size">${size} items</div>
                <div class="pack-quantity">${quantity} ${quantity === 1 ? 'pack' : 'packs'}</div>
            `;
            breakdown.appendChild(packDiv);

            totalItems += parseInt(size) * quantity;
            totalPacks += quantity;
        });

        const overshipping = totalItems - orderAmount;
        summary.innerHTML = `
            <strong>Summary:</strong><br>
            Total items shipped: ${totalItems}<br>
            Total packs: ${totalPacks}<br>
            ${overshipping > 0 ? `Overshipping: ${overshipping} items` : 'Exact match!'}
        `;

        results.classList.remove('hidden');
    }

    showError(message) {
        const errorDiv = document.getElementById('error');
        errorDiv.textContent = message;
        errorDiv.classList.remove('hidden');
    }

    showLoading() {
        document.getElementById('loading').classList.remove('hidden');
    }

    hideLoading() {
        document.getElementById('loading').classList.add('hidden');
    }

    hideAllMessages() {
        document.getElementById('results').classList.add('hidden');
        document.getElementById('error').classList.add('hidden');
        document.getElementById('loading').classList.add('hidden');
    }
}

// Start when page loads
document.addEventListener('DOMContentLoaded', () => {
    new PackCalculator();
}); 