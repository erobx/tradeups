import { useCallback, useEffect, useRef, useState } from "react"

const imgWidth = 150
const imgHeight = 150

function ItemCarousel() {
  const [activeItem, setActiveItem] = useState(0)

  const carouselImgSrc = [
    "/m4a4-howl.png",
    "/de-printstream.png",
    "/sg-tornado.png",
    "/m4a1-black-lotus.png",
    "/fs-candy-apple.png",
    "/aug-wings.png",
    "/ak-slate.png",
    "/ak-slate.png",
    "/ak-slate.png",
    "/ak-slate.png",
]

  const carouselRef = useRef(null)

  const scrollItem = () => {
    setActiveItem(prevState => {
      if (carouselImgSrc.length-1 > prevState) {
        return prevState + 1
      } else {
        return 0
      }
    })
  }

  const autoplay = useCallback(() => setInterval(scrollItem, 2000),[])

  useEffect(() => {
    const play = autoplay()
    return () => clearInterval(play)
  }, [autoplay])

  useEffect(() => {
    carouselRef.current?.scroll({left: imgWidth * activeItem})
  }, [activeItem])

  return (
    <div ref={carouselRef} className="carousel carousel-center rounded-box w-1/2">
      {
        carouselImgSrc.map((imgSrc, index) =>
          <div key={index} className="carousel-item">
            <img
              width={imgWidth}
              height={imgHeight}
              src={imgSrc}
              alt={imgSrc}
            />
          </div>
        )
      }
    </div>
  )
}

export default ItemCarousel
