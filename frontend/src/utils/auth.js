import store from "@/store";

export const getToken = function getToken () {
  return store.state.auth.token
}

