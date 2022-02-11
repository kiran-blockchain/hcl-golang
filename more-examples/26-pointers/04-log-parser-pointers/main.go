// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"strconv"
)

func main() {
	p := newParser()

	in := bufio.NewScanner(os.Stdin)
	for in.Scan() {
		parsed := parse(&p, in.Text())
		update(&p, parsed)
	}

	summarize(p)
	dumpErrs([]error{in.Err(), p.lerr})
}

// summarize summarizes and prints the parsing result
func summarize(p parser) {
	sort.Strings(p.domains)

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	for _, domain := range p.domains {
		parsed := p.sum[domain]
		fmt.Printf("%-30s %10d\n", domain, parsed.visits)
	}

	// Print the total visits for all domains
	fmt.Printf("\n%-30s %10d\n", "TOTAL", p.total)
}

// dumpErrs simplifies handling multiple errors
func dumpErrs(errs []error) {
	for _, err := range errs {
		if err != nil {
			fmt.Println("> Err:", err)
		}
	}
}


// result stores the parsed result for a domain
type result struct {
	domain string
	visits int
	// add more metrics if needed
}

// parser keep tracks of the parsing
type parser struct {
	sum     map[string]result // metrics per domain
	domains []string          // unique domain names
	total   int               // total visits for all domains
	lines   int               // number of parsed lines (for the error messages)
	lerr    error             // the last error occurred
}

// newParser constructs, initializes and returns a new parser
func newParser() parser {
	return parser{sum: make(map[string]result)}
}

// parse parses a log line and returns the parsed result with an error
func parse(p *parser, line string) (parsed result) {
	if p.lerr != nil {
		return
	}

	p.lines++

	fields := strings.Fields(line)
	if len(fields) != 2 {
		p.lerr = fmt.Errorf("wrong input: %v (line #%d)", fields, p.lines)
		return
	}

	parsed.domain = fields[0]

	var err error

	parsed.visits, err = strconv.Atoi(fields[1])
	if parsed.visits < 0 || err != nil {
		p.lerr = fmt.Errorf("wrong input: %q (line #%d)", fields[1], p.lines)
	}
	return
}

// update updates all the parsing results using the given parsing result
func update(p *parser, parsed result) {
	if p.lerr != nil {
		return
	}

	domain, visits := parsed.domain, parsed.visits

	// Collect the unique domains
	if _, ok := p.sum[domain]; !ok {
		p.domains = append(p.domains, domain)
	}

	// Keep track of total and per domain visits
	p.total += visits

	// create and assign a new copy of `visit`
	p.sum[domain] = result{
		domain: domain,
		visits: visits + p.sum[domain].visits,
	}
}
