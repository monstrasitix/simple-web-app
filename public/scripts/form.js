// @ts-check

/**
 * @param {SubmitEvent} event
 */
function onSubmit(event) {
	event.preventDefault();

	if (event.target instanceof HTMLFormElement) {
		console.log(event.target.elements);
	}
}


/**
 * @param {Event} event
 */
function onChange(event) {
	if (event.target instanceof HTMLInputElement) {
		console.log(event.target.form);
	}
}


document.addEventListener('submit', onSubmit);
document.addEventListener('change', onChange);
