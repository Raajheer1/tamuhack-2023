<template>
  <div class="mt-32 w-full p-4">  
    <div class="text-center">
      <h1 class="font-bold text-4xl">Welcome to TimelyTravel</h1>
    </div>

    <h1 class="font-bold text-sm mt-20">Full Name</h1>
    <div class="flow-root mt-2">  
      <div>
        <input name="go" v-model="name" type="text" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
      </div>
    </div>

    <h1 class="font-bold text-sm mt-6">Email</h1>
    <div class="flow-root mt-2">  
      <div>
        <input name="go" v-model="email" type="email" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
      </div>
    </div>

    <h1 class="font-bold text-sm mt-6">Password</h1>
    <div class="flow-root mt-2">  
      <div>
        <input name="go" v-model="password" type="password" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
      </div>
    </div>
    
    <button type="button" @click="register" class="mb-2 w-full inline-block px-6 py-2.5 bg-blue-600 mt-8 text-white font-medium leading-normal uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Sign Up</button>

    <button type="button" @click="router.push('/login')" class="mb-2 w-full inline-block px-6 py-2.5 mt-8 text-black font-medium">I have an account</button>
  </div>
</template>

<script setup lang="ts">
import {useRouter} from "vue-router";
import {ref} from "vue";
import axios from "axios";
const router = useRouter();

const name = ref("");
const email = ref("");
const password = ref("")
const error = ref("");
const request = {
  name,
  email,
  password
}

async function register() {
  const response = await axios.post(`${import.meta.env.VITE_API_URL}/auth/register`, request);
  if(response.status == 200){
    await router.push('/')
  }else{
    error.value = response.data.message;
  }
}

</script>

<style scoped>

</style>