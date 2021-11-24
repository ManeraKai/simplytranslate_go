// Dynamic textarea height.
const textareas = document.getElementsByTagName("textarea")
for (const i in textareas) {
    if (Object.hasOwnProperty.call(textareas, i)) {
        const textarea = textareas[i];
        textarea.style.height = "";
        textarea.style.height = textarea.scrollHeight + 10 + "px";
        function resize() {
            textarea.style.height = "";
            textarea.style.height = textarea.scrollHeight + 10 + "px";
        }
        textarea.addEventListener("input", resize);
        textarea.addEventListener("resize", resize)
    }
}
document.onkeydown =
    function (e) {
        if (e.ctrlKey && (e.code == 'Enter' || e.code == 'NumpadEnter'))
            document.getElementById('form').submit();
    }