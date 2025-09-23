// Update data-count
document.addEventListener('input', (e) => {
    if (e.target.matches('.input-container input, .input-container textarea')) {
        e.target
            .closest('.input-container')
            .setAttribute('data-count', e.target.value.length);
    }
});
