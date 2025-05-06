async function register(event: Event) {
    event.preventDefault()

    let formData = new FormData(event.target as HTMLFormElement);
    let jsonData: Record<string, string> = {};

    formData.forEach((value, key) => jsonData[key] = value as string);

    try {
        const response = await fetch("/register_submit", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(jsonData),
        });

        if (!response.ok) throw new Error(`Ошибка HTTP: ${response.status}`);

        const data = await response.json();
        console.log("Ответ сервера:", data);
    } catch (error) {
        console.error("Ошибка:", error);
    }
}