<template>
  <div class="bg-slate-100 w-full mx-auto mt-8 p-4">  
    <h1 class="font-bold">Pay for your flight</h1>

    <h1 class="font-bold text-sm mt-6">Card Number</h1>
    <div class="flow-root mt-2">  
      <div>
        <input name="go" type="text" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="1234 5678 9012 3456">
      </div>
    </div>

    <div class="flow-root mt-8">  
      <div class="float-left w-[48%]">
        <h1 class="font-bold text-sm mb-2">Exp. Date</h1>
        <input name="go" type="text" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="MM/YY">
      </div>
      <div class="float-right w-[48%]">
        <h1 class="font-bold text-sm mb-2">CVC</h1>
        <input name="back" type="text" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="123">
      </div>
    </div>

    <h1 class="font-bold text-sm mt-6">Name on Card</h1>
    <div class="flow-root mt-2">  
      <div>
        <input name="go" type="text" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500">
      </div>
    </div>
    
    <button type="button" class="mb-2 w-full inline-block px-6 py-2.5 bg-blue-600 mt-8 text-white font-medium text-xs leading-normal uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Pay</button>
  </div>
</template>

<script setup lang="ts">
  import { ref, onMounted, computed } from 'vue'
  import axios from 'axios'

</script>

<style scoped>

</style>