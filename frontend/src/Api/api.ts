


export class Api {
    url: string = "http://localhost:9090";
    proxy: Proxy = new Proxy(this.url);
}

class Proxy {

    url:string
    constructor(url:string){
        this.url=url
    }

    async GetApps(): Promise<{ success: boolean; response: any }> {
        const route = "/get-apps";

        console.log("Proxy Request:", { route });

        const res = await fetch(`${this.url}${route}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            }
        });

        const response = await res.json();

        const result = {
            success: res.ok,
            response
        };

        console.log("Proxy Response:", {
            route,
            status: res.status,
            ...result
        });

        return result;
    }

    async CreateApp(app_name: string, url: string): Promise<{ success: boolean; response: any }> {
        const route = "/register-app";

        const body = {
            name: app_name,
            url: url
        };

        console.log("Proxy Request:", { route, params: body });

        const res = await fetch(`${this.url}${route}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body)
        });

        const response = await res.json();

        const result = {
            success: res.ok,
            response
        };

        console.log("Proxy Response:", { route, status: res.status, ...result });

        return result;
    }

    async UpdateApp(app_name: string, newName: string, url: string): Promise<{ success: boolean; response: any }> {
        const route = "/update-app";

        const body = {
            name: app_name,
            newName: newName,
            url: url
        };

        console.log("Proxy Request:", { route, params: body });

        const res = await fetch(`${this.url}${route}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body)
        });

        const response = await res.json();

        const result = {
            success: res.ok,
            response
        };

        console.log("Proxy Response:", {
            route,
            status: res.status,
            ...result
        });

        return result;
    }

    async DeleteApp(id: number): Promise<{ success: boolean; response: any }> {
        const route = "/delete-app";

        const body = { id };

        console.log("Proxy Request:", { route, params: body });

        const res = await fetch(`${this.url}${route}`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(body)
        });

        const response = await res.json();

        const result = {
            success: res.ok,
            response
        };

        console.log("Proxy Response:", {
            route,
            status: res.status,
            ...result
        });

        return result;
    }
}

export const api = new Api();