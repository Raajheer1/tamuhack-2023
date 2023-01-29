<template>


  <div class="w-full mx-auto p-4">
    <h1 class="font-bold">Search Flights</h1>
    <div class="flex mt-4 gap-x-10">
      <div class="form-check form-check-inline">
        <input type="radio" v-model="picked" name="inlineRadioOptions" id="inlineRadio1" value="true" class="peer form-check-input form-check-input appearance-none rounded-full h-4 w-4 border border-gray-300 bg-white checked:bg-sky-300 checked:border-sky-300 focus:outline-none transition duration-200 mt-1.5 align-top bg-no-repeat bg-center bg-contain float-left mr-1 cursor-pointer" />
        <label for="inlineRadio1" class="w-10 h-10 bg-gray-300 transition duration-200 border-sky-300 px-2 py-1 rounded-full peer-checked:bg-sky-300 text-xs">One Way</label>
      </div>
      <div class="form-check form-check-inline">
        <input type="radio" v-model="picked" name="inlineRadioOptions" id="inlineRadio2" value="false" class="peer ml-4 form-check-input form-check-input appearance-none rounded-full h-4 w-4 border border-gray-300 bg-white checked:bg-sky-300 checked:border-sky-300 focus:outline-none transition duration-200 mt-1.5 align-top bg-no-repeat bg-center bg-contain float-left mr-1 cursor-pointer" />
        <label for="inlineRadio2" class="w-10 h-10 bg-gray-300 transition duration-200 px-2 py-1 rounded-full peer-checked:bg-sky-300 text-xs">Round Trip</label>
      </div>
    </div>

    <h1 class="font-bold text-sm mt-6">Travel Dates</h1>
    <div class="flow-root">
      <div class="float-left w-[48%]">
        <label class="text-xs mt-0">Departure</label>
        <input name="go" v-model="departure_date" type="text" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="MM/DD/YYYY">
      </div>
      <div v-if="picked !== 'true'" class="float-right w-[48%]">
        <label class="text-xs mt-0">Return</label>
        <input name="back" v-model="return_date" type="text" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="MM/DD/YYYY">
      </div>
    </div>

    <h1 class="font-bold text-sm mt-8">Locations</h1>
    <div class="flow-root">
      <div class="float-left w-[48%]">
        <label class="text-xs mt-0">Departure</label>
        <input name="start" type="text" v-model="dAirport" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="New York">
        <ul v-if="dAirport !== ''" class="p-2 rounded border-gray-300">
          <li v-for="result in dAList" :key="result.code" @click="setDep(result.code)" class="mt-2 w-full">{{ result.city }} ({{ result.code }})</li>
        </ul>
      </div>
      <div class="float-right w-[48%]">
        <label class="text-xs mt-0">Return</label>
        <input name="start" type="text" v-model="aAirport" class="top-0 bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Los Angeles">
        <ul v-if="aAirport !== ''" class="p-2 rounded border-gray-300">
          <li v-for="result in aAList" :key="result.code" @click="setArr(result.code)" class="mt-2 w-full">{{ result.city }} ({{ result.code }})</li>
        </ul>
      </div>
    </div>
    
    <button type="button" @click="search" class="mb-2 w-full inline-block px-6 py-2.5 bg-blue-600 mt-8 text-white font-medium text-xs leading-normal uppercase rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out">Search</button>
  </div>
</template>

<script setup lang="ts">
  import { ref, onMounted, computed } from 'vue'
  import {Airport} from "@/types";
  import axios from 'axios'
  import {SearchRequest} from "@/types";
  import useSearchStore from "@/store/search";
  import {Flight} from "@/types";

  const searchStore = useSearchStore();
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
    return airports.value.filter((airport: Airport) => airport.city.includes(dAirport.value) || airport.code.includes(dAirport.value))
  })

  const aAList = computed(() => {
    if(airports.value == null) {
      return []
    }
    return airports.value.filter((airport: Airport) => airport.city.includes(aAirport.value) || airport.code.includes(aAirport.value))
  })

  const departure_date = ref('')
  const return_date = ref('')

  async function search() {
    let request = {
      "departure_day": departure_date.value.substring(6, 10) + "-" + departure_date.value.substring(0, 2) + "-" + departure_date.value.substring(3, 5),
      "flight_destination": aAirport.value,
      "flight_origin": dAirport.value,
      "return_day": ""
    }


    if (picked.value !== 'true') {
      request["return_day"] = return_date.value
    }

    // await searchStore.search(request)

    const planes:Flight[] = []
    for(var i = 0; i < 10; i++){
      planes.push(generateFlight(dAirport.value, aAirport.value))
    }

    searchStore.setQuery(dAirport.value, aAirport.value, planes)

  }

function generateFlight(dep: string, arr: string) : Flight {
  const number = Math.floor(Math.random() * 1000)
  const departure:Airport = airports.value.find((airport: Airport) => airport.code === dep)
  const arrival:Airport = airports.value.find((airport: Airport) => airport.code === arr)

  let dep_time = Math.floor(Math.random() * 12)+1
  const min_interval = Math.floor(Math.random() * 4)
  let mins = "00"
  if(min_interval == 0){
    mins = "15"
  }else if(min_interval == 1) {
    mins = "30"
  }else if(min_interval == 2) {
    mins = "45"
  }else if(min_interval == 3) {
    mins = "25"
  }

  let mins_offset = Math.floor(Math.random() * 3) * 10

  let total_min = parseInt(mins) + mins_offset
  if(total_min > 60){
    total_min = total_min - 60
    dep_time = dep_time + 1
  }

  let month = ""
  if(departure_date.value.substring(0, 2) == "01"){
    month = "Jan"
  }else if(departure_date.value.substring(0, 2) == "02") {
    month = "Feb"
  }else if(departure_date.value.substring(0, 2) == "03") {
    month = "Mar"
  }else if(departure_date.value.substring(0, 2) == "04") {
    month = "Apr"
  }else if(departure_date.value.substring(0, 2) == "05") {
    month = "May"
  }else if(departure_date.value.substring(0, 2) == "06") {
    month = "Jun"
  }else if(departure_date.value.substring(0, 2) == "07") {
    month = "Jul"
  }else if(departure_date.value.substring(0, 2) == "08") {
    month = "Aug"
  }else if(departure_date.value.substring(0, 2) == "09") {
    month = "Sep"
  }else if(departure_date.value.substring(0, 2) == "10") {
    month = "Oct"
  }else if(departure_date.value.substring(0, 2) == "11") {
    month = "Nov"
  }else if(departure_date.value.substring(0, 2) == "12") {
    month = "Dec"
  }


  const delay_perc = Math.floor(Math.random() * 25)
  const delay_len = Math.floor(Math.random() * 60)

  const plane = {
    flight_number: `AA${number}`,
    departure: departure,
    arrival: arrival,
    date: `${departure_date.value.substring(3, 5)}${month}`,
    departure_time: `${dep_time}:${mins}`,
    arrival_time: `${dep_time+3}:${total_min}`,
    delay_percentage: delay_perc,
    delay_time: delay_len,
    price: number/4
  };

  return plane
}

</script>

<style scoped>

</style>