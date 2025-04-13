## Specialty Coffee Crawler
The Specialty Coffee Finder is a tool designed to connect coffee
enthusiasts with high-quality specialty coffee sources by aggregating
publicly available information from coffee roaster websites.
Our aim is to create a centralized resource that promotes specialty cofee
businesses and enhances discoverability for consumers.


---- To-Do List

Phase 1: Planning
1. Define MVP Scope
* Decide core features (e.g., search, roaster profiles, basic listings)
* Write a short document or list outlining what's in the MVP vs future enhancements.
2. Interact/Ethical Guidelines
* Review web scraping laws and best practices (e.g., respect robots.txt, tos).
* Reach out to then initial 50 prospect roasters and ask for writtern permission.

Phase 2: Frontend Development
3. Build Basic Frontend
* Set up Htmx and Templ in the web( in backend for now) folder.
* Add Tailwind Css for styling.
* Create a homepage with a search bar and placeholder results.
4. Add Core UI Components
* Implement "Search UI" (text input with submit button).
* Add "Filters" (Dropdown for roast level, origin, etc., harcoded for now).
* Create a "Roaster Profiles" template (e.g., card with name, description, link).
5. Test Frontend with Dummy Data
* Hardcode sample roaster data (e.g., 3-5 entries) to display in the UI.
* Ensure responsive design works on mobile and desktop.

Pashe 3: Backend Development
6. Set up API Server
* Initialize a Golang project.
* Create a simple endpoint (e.g., /api/roasters) that returns dummy JSON data.
* Connect Frontend to API using htmx fetch(basic js).
7. Build Web Scraper
* Set up a Golang or Python Script with Colly or BeautifulSoup.
* Scrape basic data (e.g., roaster name, location, website) from 1-2 coffee-sites.
* Output scraped data to JSON file.
8. Integrate Scraper with Backend
* Modify scraper to send data to the API Server (e.g., POST request).
* Add an endpoint to receive and store scraped data temporarily (in memory or file.)

Phase 4: Data Storage
9. Set up Database
* Ininitalize PostgreSQL in docker
* Create a table for roasters (this should be done when backend server starts).
* Insert sample data manually to test.
10. Connect Backend to Database
* Update /api/roasters endpoint to fetch data from the database.
11. Add Search Functionality
* Install Elasticsearch(variation) locally.
* Index roaster data from database into Elasticsearch.
* Create an API endpoint (e.g., /api/search) to query Elasticsearch.

Phase 5: Integration and Polish
12. Add Map View if we implement maps
* Sign up for Google Maps API and get an API Key.
* Integate a map component in the Frontend.
* Plot sample roaster locations from the database.
13. Test End-to-End Flow
* Scrape data -> store in database -> search via Frontend -> display results.
* Fix any bugs or UI issues.
14. Add External Integrations (Optional for MVP)
* Pull X posts about coffee roasters using the X API (requires API access).
* Set up SendGrid(or SendSpectra) for basic email notifications (e.g., "New roasters added").
Phase 6: Deployment and Iteration 
15. Deploy the Application
* Containerize Backend and Frontend with Docker/Podman.
* Deploy to a cloud provider (e.g., Hetzner cloud).
* Set up a simple domain or use a free subdomain.
16. Gather Feedback
* Share the mvp with small group (e.g., reddit/r/sideprojects, r/espresso, etc).
* Collect feedback on usability, features, and bugs.
17. Plan Next Features
* Add caching with Redis for faster searches.
* Expand scraper to more websiter.
* Implement user accounts  or analytics for roasters.

Ongoing Tasks
18. Monitor and Maintain
* Check scraper for errors or blocked sites.
* Update database with new roasters periodically.
19. Document Progress
* Keep a log of completed tasks and issues in your Git repo (e.g., Completed.md).

    
