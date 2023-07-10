/**
 * @param {SubmitEvent} event
 */
function onSubmit({ target }) {
	if (target instanceof HTMLFormElement) {
		const data = new FormData(target);
		const output = {};

		for (const [name, value] of data.entries()) {
			output[name] = value;
		}

		console.log(output);
	}
}

/**
 * @param {Event} event
 */
function onSubmit({ target }) {
	if (target instanceof HTMLInputElement) {
		const data = new FormData(target.form);

		console.log({
			[target.name]: data.get(target.name),
		});
	}
}

document.addEventListener("submit", onSubmit);
document.addEventListener("change", onChange);
