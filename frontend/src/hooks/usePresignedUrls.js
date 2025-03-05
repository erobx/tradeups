import { useState, useEffect } from "react"

const URL_CACHE_KEY = "presigned_url_cache"

export function getCache() {
    const cache = localStorage.getItem(URL_CACHE_KEY)
    return cache ? JSON.parse(cache) : {}
}

export function setCache(cache) {
    localStorage.setItem(URL_CACHE_KEY, JSON.stringify(cache))
}

/*
 * @returns {string[]}
 * 
 */
export function usePresignedUrls(inventory) {
    const [processedInventory, setProcessedInventory] = useState(inventory)

    useEffect(() => {
        const cache = getCache()
        const now = Date.now()
        const newCache = {}

        const updatedInventory = inventory.map(item => {
            const cachedUrl = cache[item.imageSrc]

            if (cachedUrl && cachedUrl.expiry > now) {
                newCache[item.imageSrc] = cachedUrl
                return { ...item, imageSrc: cachedUrl.url }
            }

            newCache[item.imageSrc] = {
                url: item.imageSrc,
                expiry: now + (23 * 60 * 60 * 1000)
            }

            return item
        })    

        setCache(newCache)
        setProcessedInventory(updatedInventory)
    }, [inventory])

    return processedInventory
}
