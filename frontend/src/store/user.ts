import { defineStore } from "pinia";

import { API } from "@/utils/api";

interface UserState {
    user: User | null;
    fetching: boolean;
    hasFetched: boolean;
    loading: Promise<void> | null;
}

interface User {
    id: number;
    first_name: string;
    last_name: string;
    ffnumber: string;
    created_at: string;
    updated_at: string;
}


const useUserStore = defineStore("user", {
    state: () =>
        ({
            user: null,
            fetching: false,
            hasFetched: false,
        } as UserState),
    getters: {
        isLoggedIn: (state) => !!state.user,
    },
    actions: {
        async fetchUser() {
            this.fetching = true;
            try {
                const { data } = await API.get("/v1/user/");
                this.user = data;
            } catch (e) {
                this.user = null;
            } finally {
                this.fetching = false;
                this.hasFetched = true;
            }
        },
    },
});

export default useUserStore;