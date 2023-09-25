<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<div align="center">
<br />
  <a href="https://github.com/christian-jaimes/pokemon-data-scraping">
    <img src="image/project-logo.png" alt="Logo">
  </a>

  <h1 align="center">Pokemon data scraping</h1>

  <p align="center">
    Collecting comprehensive Pokemon data, from their abilities and stats to their name etymology. It's a one-stop resource for Pokemon enthusiasts
    <br />
    <br />
    <a href="https://github.com/christian-jaimes/pokemon-data-scraping">View Demo</a>
    -
    <a href="https://github.com/christian-jaimes/pokemon-data-scraping/issues">Report Bug</a>
    -
    <a href="https://github.com/christian-jaimes/pokemon-data-scraping/issues">Request Feature</a>
  </p>
</div>

## About The Project

The **Pokemon Data Scraping** project is a comprehensive data collection effort aimed at gathering valuable information about Pokemon creatures. Through web scraping techniques, this project extracts and organizes data from the PokemonDB to create a detailed dataset of Pokemon attributes and characteristics.

### Project Scope

The project offers two main data collection modes:
1. **Full Dataset:** Scrapes data for all available Pokemon species, providing an extensive dataset that covers the entire Pokemon universe.
2. **Partial Dataset:** Allows for focused scraping, which is useful for testing.

### Data Output

The gathered data is structured into five tables which are interconnected:
<div align="center">
<br />
<img src="docs/model/Pokemon data model.png" alt="pokemon data model">
</div>

- **Pokemon Details:** Includes essential information such as Pokemon ID, name, description, image, species, height, weight, types, abilities, EV yield, catch rate, base experience, growth rate, gender distribution, generation, and name etymology.
- **Pokemon Stats:** Provides detailed statistics on HP, Attack, Defense, Special Attack, Special Defense, and Speed. This information is valuable for analyzing the combat capabilities of each Pokemon.
- **Pokemon League:** Collects data about Pokemon leagues and their associated regions, offering insights into the competitive Pokemon gaming scene.
- **Pokemon region details:** Contains detailed information on the regions like, professors, villians and starter pokemons.
- **Pokemon region coordinates:** This is a manual dataset made with [cbistudio](https://cbistudio.interworks.com/) to calculate the X and Y Coordinates to plot the regions into a custom made map.  
<div align="center">
<br />
<img src="docs/region-map.png" alt="Region map">
</div>

### Why This Project?

This project serves as a valuable resource for Pokemon enthusiasts, researchers, and developers who require comprehensive and structured Pokemon data. Whether you're building a Pokemon-related application, conducting research, or simply exploring the vast world of Pokemon, this dataset can save you hours of manual data collection and organization.

This project is one of many I made to master python for data, others can be found in my [portfolio](https://https://christian-jaimes.github.io/).

Contributions and suggestions from the community to enhance and expand this project are always welcome. Feel free to fork the repository, create pull requests, or open issues to help us improve this valuable resource.

### Built With

[![Python][python]][python-url]
[![Visual studio code][vscode]][vscode-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>
    
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

Before you begin, ensure you have the following prerequisites installed:
- [Python](https://www.python.org/downloads/): This project is written in Python, so you'll need Python installed on your machine. You can download the latest version from the official website.


### Installation

Follow these steps to get a local copy of the project up and running:

1. Clone the repository to your local machine:

   ```sh
   git clone https://github.com/christian-jaimes/pokemon-data-scraping.git
   

2. Install necessary the libraries:

   ```sh
   pip install pandas
   pip install beautifulsoup4

3. You can find the Jupyter notebook file in the **notebooks** folder

4. Have fun 😁

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

## Contact

[![Github][github]][github-url]
[![LinkedIn][linkedin-shield]][linkedin-url]


## Credits

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [Choose an Open Source License](https://choosealicense.com)
* [Img Shields](https://shields.io)
* [Simple Icons](https://simpleicons.org/)
* [GitHub Pages](https://pages.github.com)
* [Pokemon DB](https://pokemondb.net/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[contributors-shield]: https://img.shields.io/github/contributors/christian-jaimes/pokemon-data-scraping.svg?style=for-the-badge
[contributors-url]: https://github.com/christian-jaimes/pokemon-data-scraping/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/christian-jaimes/pokemon-data-scraping.svg?style=for-the-badge
[forks-url]: https://github.com/christian-jaimes/pokemon-data-scraping/network/members

[stars-shield]: https://img.shields.io/github/stars/christian-jaimes/pokemon-data-scraping.svg?style=for-the-badge
[stars-url]: https://github.com/christian-jaimes/pokemon-data-scraping/stargazers

[issues-shield]: https://img.shields.io/github/issues/christian-jaimes/pokemon-data-scraping.svg?style=for-the-badge
[issues-url]: https://github.com/christian-jaimes/pokemon-data-scraping/issues

[license-shield]: https://img.shields.io/github/license/christian-jaimes/pokemon-data-scraping.svg?style=for-the-badge
[license-url]: https://github.com/christian-jaimes/pokemon-data-scraping/blob/master/LICENSE.txt



[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/christian-jaimes

[python]: https://img.shields.io/badge/Python-20232A?style=for-the-badge&logo=python&logoColor=4584B6
[python-url]: https://www.python.org/

[vscode]: https://img.shields.io/badge/VSCode-20232A?style=for-the-badge&logo=visualstudiocode&logoColor=0078d7
[vscode-url]: https://code.visualstudio.com/

[github]: https://img.shields.io/badge/github-20232A?style=for-the-badge&logo=github&logoColor=ffffff
[github-url]: https://github.com/christian-jaimes

[tableau]: https://img.shields.io/badge/tableau-20232A?style=for-the-badge&logo=tableau&logoColor=E97627
[tableau-url]: https://public.tableau.com/app/profile/christian-jaimes
