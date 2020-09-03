# Methods
Each language's buggy code has two bugs. One is a compilation bug to test how easy it is to follow the compiler's error message and the second is a logical bug to check how easy it is to grasp and understand the syntax of the language.

For any type of analysis to happen, we have to first collect data!

We obtained various data points during the experiment based on which we made various conclusions which are given in the conclusions section. 

<!--
Participants in the experiment will give us the following points:

Languages the participant is familiar with. How familiar the participant is with Go, Perl and Ruby. Rank the languages based on how easy it was to debug among the three languages. While the participants are debuging the code in each language we will gather the following data:

Debug time taken in each language. Number of times the participant ran code for each language. Based on the data collected we will compare each language and draw our conclusions!
-->

The data points given by participants using a Survey Monkey form are:

- Languages the participant is familiar with.
- How familiar the participant is with each language.
- How similar each language looked to the languages participants already know.
- Rank of languages based on how easy it was to debug among the three languages.

The data points collected by the team during thr experiment are:

- Debug time taken in each language.
- Number of times the participant ran code for each language.

Based on the data collected we compared the languages on following criterias:

- Overall debug time for each language.
- Similarity to other languages.
- Participants ranking of languages for easiness in debugging.

# Materials  
### Survey
We used Survey Monkey as a questionnaire to collect data from the participants. 

### Code with bugs  
Our task sets two bugs for each language, one is compile bug, and another is logic bug. Each language has bugs designed in the same place. So after debugging one of the languages the participant might spend less time finding the bug's location while debugging the other two languages. Therefore, we changed the order of languages for each participant. 

### Online ide  
Partcipants were provided with an online ide link 'https://www.onlinegdb.com/' and required to share their screen during the experienments.

### Data collected  
By monitoring the parcipants debugging process, we recorded how long it took to debug for each different language and evaluated if they passed the test. We also counted how many times participants excute the code in order to see if the warning messages gave enough cue for participants to find the bug right away.
At the end of the questionnaire, there is a question ask "how similar GO/Perl/Ruby looked as compared to the languages you have worked with". This is designed to consider whether it affects the debug speed of the participants.

# Observations

### Plot to describe percentage of participants familiar with each language

![Plot of languages known by participants](https://github.com/Zhuolin0212/GOL_Group_10/blob/master/Plots/Languages_Known.JPG)

### Plots of similarity to already known languages

#### Plot of how similar Go looked as compared to languages participants are already familiar with

![Plot of languages known by participants](https://github.com/Zhuolin0212/GOL_Group_10/blob/master/Plots/Similarity_To_Go.JPG)

#### Plot of how similar Ruby looked as compared to languages participants are already familiar with

![Plot of languages known by participants](https://github.com/Zhuolin0212/GOL_Group_10/blob/master/Plots/Similarity_To_Ruby.JPG)

#### Plot of how similar Perl looked as compared to languages participants are already familiar with

![Plot of languages known by participants](https://github.com/Zhuolin0212/GOL_Group_10/blob/master/Plots/Similarity_To_Perl.JPG)

### Plots of debug time in each language

#### Histogram of debug time for participants during the experiment in Go

![Plot of languages known by participants](https://github.com/Zhuolin0212/GOL_Group_10/blob/master/Plots/GO_Debug_Histogram.JPG)

#### Histogram of debug time for participants during the experiment in Ruby

![Plot of languages known by participants](https://github.com/Zhuolin0212/GOL_Group_10/blob/master/Plots/Ruby_Debug_Histogram.JPG)

#### Histogram of debug time for participants during the experiment in Perl

![Plot of languages known by participants](https://github.com/Zhuolin0212/GOL_Group_10/blob/master/Plots/Perl_Debug_Histogram.JPG)

### Difficulty in debugging ranks for languages by participants

![Plot of languages known by participants](https://github.com/Zhuolin0212/GOL_Group_10/blob/master/Plots/Languages_Difficulty_Rank.JPG)

```
Based on the observations that we collected from 9 participants, we can deduce following points:
- Rank from overall lowest time taken to overall highest time taken for debugging: 1) Go 2) Perl 3) Ruby 
- Rank from most similar to least similar languages known by participants: 1) Go 2) Ruby 3) Perl
- Rank from easier to debug to hard to debug  based on participants answers: 1) Go 2) Tie between Ruby & Perl
```

# Conclusions

- Among the three languages Go is undoubtedly the easiest to debug in as participants took the lowest time to debug, participants found it more similar to previously known languages & participants found it to be the most easy language to debug in among the three languages.
- Perl comes in at second place as it came second in the ranking for debug time, tied with Ruby in the ranking by participants for easiness in debugging and came in third in the ranking for similarity to previously know languages by participants.
- Ruby comes in at third place  as it came third in the ranking for debug time, tied with Ruby in the ranking by participants for easiness in debugging and came in second in the ranking for similarity to previously know languages by participants.

(We have given more importance to debug time taken by participant than similarity to languages already known by participants)

<!--In this experiment, we invited 9 participants to debug the bug version of Game of Life in 3 different languages(Ruby, Go, Perl). Participants performed a 30-minute test, with about 8 minutes of debug time for each language and a post-questionnaire. The questionnaire showed that Ruby was the most difficult language for participants to debug, followed by Perl, and the Go language that was the easiest to debug. The most familiar languages for the participants themselves are JS, C, Java and Python. And the Go language is the most similar to what they mastered at before, Perl is completely different from what they were good at before and Ruby is somewhat similar. Ruby is similar to what they learned before, but the debug success rate is low. We think it may be that the bug needs to be modified slightly different from the other 2 languages. 
-->
# Threats to validity  

The above conclusions are subject to the validity of the experiments that we conducted. Some of the threats to validity that we found are:  
- Insufficient data points, since we could only conduct experiment with 9 different participants. Also, we think even 10 participants would not have been sufficient to draw reliable and trusty conclusions. Next time, we can increase the size of data points helping us draw reliable conclusion.
- Maladroit calculation of time taken to debug the codes as we calculated the time manually rather than automating it. So, we will automate the calculation of debug time.
- Ineffectual bugs in the 3 codes, since we had the same exact compilation and logical bugs in all the 3 codes. So next time, we can have different bugs in different codes keeping the difficulty level of bugs in each language same.
- Inconsistent level of complexity of the bugs in all the 3 languages. Participants found the logical bug of Ruby to be quite difficult because of the style of writing. Here, we can keep the level of complexity for all the languages to be the same next time.
