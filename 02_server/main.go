package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo") Formulario Navideño
}

func main() {
	// Servir archivos estáticos desde el directorio "static"
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/", helloHandler)

	fmt.Println("Servidor iniciado en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}  <!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta http-equiv="X-UA-Compatible" content="IE=edge" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />

    <!-- @@@@@@@@@@@@@@@@@@@ favicon @@@@@@@@@@@@@@@@@@@@ -->
    <link
      rel="shortcut icon"
      href="./assets/img/favicon.png"
      type="image/x-icon"
    />

    <!-- @@@@@@@@@@@@@@@@@@@ boxicons @@@@@@@@@@@@@@@@@@@@ -->
    <link
      href="https://unpkg.com/boxicons@2.1.2/css/boxicons.min.css"
      rel="stylesheet"
    />

    <!-- @@@@@@@@@@@@@@@@@@@ swiperjs.com @@@@@@@@@@@@@@@@@@@@ -->
    <link
      rel="stylesheet"
      href="https://unpkg.com/swiper@8/swiper-bundle.min.css"
    />

    <!-- @@@@@@@@@@@@@@@@@@@ CSS @@@@@@@@@@@@@@@@@@@@ -->
    <link rel="stylesheet" href="./assets/css/styles.css" />

    <!-- @@@@@@@@@@@@@@@@@@@ breakpoints CSS @@@@@@@@@@@@@@@@@@@@ -->
    <link rel="stylesheet" href="./assets/css/breakpoints.css" />

    <title>Christmas</title>
  </head>
  <body>
    <!-- @@@@@@@@@@@@@@@@@@@ pre loader @@@@@@@@@@@@@@@@@@@@ -->
    <div class="preloader" id="preloader">
      <img src="./assets/img/preloader.png" alt="" class="preloader__img" />
    </div>

    <!-- @@@@@@@@@@@@@@@@@@@ header @@@@@@@@@@@@@@@@@@@@ -->

    <header class="header" id="header">
      <div class="container">
        <nav class="nav">
          <a href="#" class="nav__logo">
            <img src="./assets/img/logo.png" alt="" /> Christmas
          </a>

          <div class="nav__menu" id="nav-menu">
            <ul class="nav__list">
              <li class="nav__item">
                <a href="#home" class="nav__link">Home</a>
              </li>
              <li class="nav__item">
                <a href="#celebrate" class="nav__link">Celebrate</a>
              </li>
              <li class="nav__item">
                <a href="#gift" class="nav__link">Gifts</a>
              </li>
              <li class="nav__item">
                <a href="#new" class="nav__link">New</a>
              </li>
            </ul>

            <div class="nav__close" id="nav-close">
              <i class="bx bx-x"></i>
            </div>

            <img src="./assets/img/nav-light.png" alt="" class="nav__corner" />
          </div>

          <div class="nav__toggle" id="nav-toggle">
            <i class="bx bx-menu-alt-right"></i>
          </div>
        </nav>
      </div>
    </header>

    <main>
      <!-- @@@@@@@@@@@@@@@@@@@ home @@@@@@@@@@@@@@@@@@@@ -->
      <section class="home" id="home">
        <div class="container">
          <div class="home__container">
            <img src="./assets/img/home.png" alt="" class="home__img" />

            <div class="home__data">
              <h1 class="home__title">
                Merry Christmas and <br />
                Happy New Year!
              </h1>
              <p class="home__description">
                Christmas and a new year is about to begin, all good wishes and
                successes.
              </p>
              <a href="#" class="btn">Get Started</a>
            </div>
          </div>
        </div>
      </section>

      <!-- @@@@@@@@@@@@@@@@@@@ giving @@@@@@@@@@@@@@@@@@@@ -->
      <section class="giving section">
        <div class="container">
          <h2 class="section-title">
            Start Giving This <br />
            Christmas
          </h2>

          <div class="giving__container">
            <div class="giving__content">
              <img src="./assets/img/giving1.png" alt="" class="giving__img" />
              <h3 class="giving__title">Surprise gifts</h3>
              <p class="giving__description">
                They are the best gifts there is.
              </p>
            </div>
            <div class="giving__content">
              <img src="./assets/img/giving2.png" alt="" class="giving__img" />
              <h3 class="giving__title">Ornaments</h3>
              <p class="giving__description">Give a moment to decorate.</p>
            </div>
            <div class="giving__content">
              <img src="./assets/img/giving3.png" alt="" class="giving__img" />
              <h3 class="giving__title">Lots of love</h3>
              <p class="giving__description">
                Give away feelings that last forever.
              </p>
            </div>
          </div>
        </div>
      </section>

      <!-- @@@@@@@@@@@@@@@@@@@ celebrate @@@@@@@@@@@@@@@@@@@@ -->
      <section class="celebrate section" id="celebrate">
        <div class="container">
          <div class="celebrate__container">
            <div class="celebrate__data">
              <h2 class="celebrate__title">
                Celebrate With A
                <br />
                Lot Of Love
              </h2>
              <p class="celebrate__description">
                In this holidays, celebrate with much love and peace, offering
                many blessings to our loved ones, friends and neighbors, wishing
                them a good Christmas message.
              </p>
              <a href="#" class="btn">Send Good Wishes</a>
            </div>

            <img
              src="./assets/img/celebrate.png"
              alt=""
              class="celebrate__img"
            />
          </div>
        </div>
      </section>

      <!-- @@@@@@@@@@@@@@@@@@@ gift @@@@@@@@@@@@@@@@@@@@ -->
      <section class="gift section" id="gift">
        <div class="container">
          <h2 class="section-title">Share A Gift</h2>

          <div class="gift__container">
            <div class="gift__card">
              <i class="bx bx-heart gift__icon"></i>
              <img src="./assets/img/gift1.png" alt="" class="gift__img" />
              <h3 class="gift__price">$15</h3>
              <span class="gift__description">Gingerbread</span>
            </div>
            <div class="gift__card">
              <i class="bx bx-heart gift__icon"></i>
              <img src="./assets/img/gift2.png" alt="" class="gift__img" />
              <h3 class="gift__price">$22</h3>
              <span class="gift__description">Santa Claus Hat</span>
            </div>
            <div class="gift__card">
              <i class="bx bx-heart gift__icon"></i>
              <img src="./assets/img/gift3.png" alt="" class="gift__img" />
              <h3 class="gift__price">$48</h3>
              <span class="gift__description">Christmas Tree</span>
            </div>
            <div class="gift__card">
              <i class="bx bx-heart gift__icon"></i>
              <img src="./assets/img/gift4.png" alt="" class="gift__img" />
              <h3 class="gift__price">$35</h3>
              <span class="gift__description">Snowman</span>
            </div>
            <div class="gift__card">
              <i class="bx bx-heart gift__icon"></i>
              <img src="./assets/img/gift5.png" alt="" class="gift__img" />
              <h3 class="gift__price">$12</h3>
              <span class="gift__description">Candy Stick</span>
            </div>
          </div>
        </div>
      </section>
	  <span class="footer__button-subtitle">GET IT ON</span>
	  <span class="footer__button-title">Google Play</span>
	</div>
  </button>
  <button class="footer__button">
	<i class="bx bxl-apple footer__button-icon"></i>
	<div>
	  <span class="footer__button-subtitle">Download on the</span>
	  <span class="footer__button-title">App Store</span>
	</div>
  </button>
</div>
</div>

<img src="./assets/img/footer1.png" alt="" class="footer__img-left" />
<img
src="./assets/img/footer2.png"
alt=""
class="footer__img-right"
/>
</div>

<p class="footer__copy">
© 2022 coded by
<a href="https://github.com/soroushmdn" target="_blank"
>Soroush Karimi.</a
>

Design from Bedimcode.
</p>
</div>
</footer>

<!-- @@@@@@@@@@@@@@@@@@@ scroll up @@@@@@@@@@@@@@@@@@@@ -->
<a href="#" class="scrollup" id="scrollup">
<i class="bx bx-up-arrow-alt"></i>
</a>

<!-- @@@@@@@@@@@@@@@@@ swiperjs.com @@@@@@@@@@@@@@@@@@ -->
<script src="https://unpkg.com/swiper@8/swiper-bundle.min.js"></script>

<!-- @@@@@@@@@@@@@@@@@ scrollrevealjs.org @@@@@@@@@@@@@@@@@@ -->
<script src="https://unpkg.com/scrollreveal"></script>

<!-- @@@@@@@@@@@@@@@@@@@ main JS @@@@@@@@@@@@@@@@@@@@ -->
<script src="./assets/js/main.js"></script>
</body>
</html>.
