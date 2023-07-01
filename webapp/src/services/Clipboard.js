async function AddTextToClipboard(text) {
    await navigator.clipboard.writeText(text);
}

export default AddTextToClipboard;

