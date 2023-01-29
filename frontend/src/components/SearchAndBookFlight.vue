<template>
  <div class="bg-slate-100 w-full mx-auto p-4">  
    <h1 class="font-bold">Search Flights</h1>
    <div class="flex mt-4">
      <div class="form-check form-check-inline">
        <input type="radio" v-model="picked" name="inlineRadioOptions" id="inlineRadio1" value="true" class="peer form-check-input form-check-input appearance-none rounded-full h-4 w-4 border border-gray-300 bg-white checked:bg-sky-300 checked:border-sky-300 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-1 cursor-pointer" />
        <label for="inlineRadio10" class="w-10 h-10 bg-gray-300 transition duration-200 border-sky-300 p-2 rounded-full peer-checked:bg-sky-300 text-xs">One Way</label>
      </div>
      <div class="form-check form-check-inline">
        <input type="radio" v-model="picked" name="inlineRadioOptions" id="inlineRadio2" value="false" class="peer ml-4 form-check-input form-check-input appearance-none rounded-full h-4 w-4 border border-gray-300 bg-white checked:bg-sky-300 checked:border-sky-300 focus:outline-none transition duration-200 mt-1 align-top bg-no-repeat bg-center bg-contain float-left mr-1 cursor-pointer" />
        <label for="inlineRadio20" class="w-10 h-10 bg-gray-300 transition duration-200 p-2 rounded-full peer-checked:bg-sky-300 text-xs">Round Trip</label>
      </div>
    </div>

    <h1 class="font-bold text-sm mt-6">Travel Dates</h1>
    <div class="flow-root mt-2">  
      <div class="float-left w-[48%]">
        <input name="go" type="text" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="MM/DD/YYYY">
      </div>
      <div v-if="picked !== 'true'" class="float-right w-[48%]">
        <input name="back" type="text" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="MM/DD/YYYY">
      </div>
    </div>

    <h1 class="font-bold text-sm mt-8">Locations</h1>
    <div class="flow-root mt-2">  
      <div class="float-left w-[48%]">
        <input name="start" type="text" v-model="dAirport" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="New York">
        <ul v-if="dAirport !== ''" class="p-2 rounded border-gray-300">
          <li v-for="result in dAList" :key="result.code" @click="setDep(result.code)" class="mt-2 w-full">{{ result.city }} ({{ result.code }})</li>
        </ul>
      </div>
      <div class="float-right w-[48%]">
        <input name="start" type="text" v-model="aAirport" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Los Angeles">
        <ul v-if="aAirport !== ''" class="p-2 rounded border-gray-300">
          <li v-for="result in aAList" :key="result.code" @click="setArr(result.code)" class="mt-2 w-full">{{ result.city }} ({{ result.code }})</li>
        </ul>
      </div>
    </div>
    <button type="button" class="mb-2 w-full inline-block px-6 py-2.5 bg-blue-600 mt-8 text-white font-medium text-xs leading-normal uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Search</button>
  </div>
</template>

<script setup lang="ts">
  import { ref, onMounted, computed } from 'vue'
  import axios from 'axios'

  const picked = ref('true')
  const dAirport = ref('')
  const aAirport = ref('')

  const airports = ref()
  onMounted(() => {
    getAirports()
  })

  async function getAirports() {
    const response = await axios.get('https://timley-travel-tamuahck.herokuapp.com/airports/all')
    airports.value = response.data
  }

  function setArr(arrival:string) {
    aAirport.value=arrival
  }

  function setDep(arrival:string) {
    dAirport.value=arrival
  }

  const dAList = computed(() => {
    if(airports.value == null) {
      return []
    }
    return airports.value.filter((airport) => airport.city.includes(dAirport.value) || airport.code.includes(dAirport.value))
  })

  const aAList = computed(() => {
    if(airports.value == null) {
      return []
    }
    return airports.value.filter((airport) => airport.city.includes(aAirport.value) || airport.code.includes(aAirport.value))
  })
</script>

<style scoped>

</style>