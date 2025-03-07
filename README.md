# prog2005
\documentclass{article}
\usepackage{hyperref}
\usepackage{listings}
\usepackage{xcolor}

\title{Country Info API}
\author{Sebastian Andersen Marterer}
\date{\today}

\begin{document}

\maketitle

\section{Introduction}
This project is a RESTful API written in Go (Golang) that provides information about countries, including population data and general country details. The API fetches data from external sources like CountriesNow and RestCountries.

\section{Installation}
\begin{enumerate}
    \item Install \textbf{Go (Golang)} if you haven't already. You can download it from:
    \begin{center}
        \url{https://golang.org/dl/}
    \end{center}
    \item Clone this repository:
    \begin{lstlisting}[language=bash]
    git clone https://github.com/yourusername/countryinfo-api.git
    cd countryinfo-api
    \end{lstlisting}
    \item Run the server:
    \begin{lstlisting}[language=bash]
    go run main.go
    \end{lstlisting}
\end{enumerate}

\section{Usage}
Once the server is running, you can access the API endpoints using \texttt{curl} or a web browser.

\subsection{Fetching Country Information}
\begin{lstlisting}[language=bash]
curl -X GET "http://localhost:8080/countryinfo/v1/info/no"
\end{lstlisting}
\textbf{Response:}
\begin{lstlisting}[language=json]
{
    "name": "Norway",
    "continents": ["Europe"],
    "population": 5379475,
    "languages": {"nno": "Norwegian Nynorsk", "nob": "Norwegian Bokmål", "smi": "Sami"},
    "borders": ["FIN", "SWE", "RUS"],
    "flag": "https://flagcdn.com/w320/no.png",
    "capital": "Oslo",
    "cities": ["Abelvaer", "Adalsbruk", "Adland", "Agotnes", "Agskardet"]
}
\end{lstlisting}

\subsection{Fetching Population Data}
\begin{lstlisting}[language=bash]
curl -X GET "http://localhost:8080/countryinfo/v1/population/no?limit=2010-2015"
\end{lstlisting}
\textbf{Response:}
\begin{lstlisting}[language=json]
{
    "mean": 5044396,
    "values": [
        {"year":2010,"value":4889252},
        {"year":2011,"value":4953088},
        {"year":2012,"value":5018573},
        {"year":2013,"value":5079623},
        {"year":2014,"value":5137232},
        {"year":2015,"value":5188607}
    ]
}
\end{lstlisting}

\subsection{Checking API Status}
\begin{lstlisting}[language=bash]
curl -X GET "http://localhost:8080/countryinfo/v1/status/"
\end{lstlisting}
\textbf{Response:}
\begin{lstlisting}[language=json]
{
    "countriesnowapi": 200,
    "restcountriesapi": 200,
    "version": "v1",
    "uptime": 3600
}
\end{lstlisting}

\section{Project Structure}
\begin{lstlisting}
countryinfo-api/
│-- main.go
│-- handlers/
│   │-- country.go
│   │-- population.go
│   │-- status.go
│-- README.tex
\end{lstlisting}

\section{Contributing}
Contributions are welcome! Please follow these steps:
\begin{enumerate}
    \item Fork the repository.
    \item Create a new branch for your feature or bug fix.
    \item Commit your changes with a descriptive message.
    \item Open a pull request for review.
\end{enumerate}

\section{License}
This project is licensed under the MIT License. See \texttt{LICENSE} for details.

\section{Contact}
For questions or issues, create a GitHub issue or reach out at:
\begin{center}
    \href{mailto:your.email@example.com}{your.email@example.com}
\end{center}

\end{document}
