#!/usr/bin/perl

use Data::Dumper;
use warnings;

my @inp_arr;
my $i = 0;

my $gen = 5;
chomp($gen);

## helper
sub getValue {
    my($x, $y ) = @_;
    my $s = 1;
    my $xbound = scalar @inp_arr;
    my $ybound = scalar @inp_arr;
        if ($x < 0 || $y < 0 || $x > $xbound-$s || $y > $ybound-$s) {
        return 0;
        } else {
            return($inp_arr[$x][$y]);
        }
}

## main
sub main {
    open (FH,'<', "input.txt") or die "Cannot open file";
    while (my $line = <FH>){
    chomp($line);
    $inp_arr[$i] = [split(",", $line)];
    $i++;
}
close(FH);

@print_arr = map { [@$_] } @inp_arr;
for (my $i = 0; $i < $gen; $i++){
    print "Generation: $i\n";
    my $e = 1;
    while($e) {
	print map((map($_ ? "X" : "o", @$_), "\n"), @inp_arr);
	$e--;
    }
    
    print "\n\n";
    my $row = @inp_arr;
    for my $j (1 .. $#inp_arr) {
        for $k (1..$#{ $inp_arr[$i] }) { 
            my $count = 0;
        for ( [$j-1, $k-1], [$j-1, $k], [$j-1, $k+1],[ $j, $k-1],[ $j, $k+1],[ $j+1, $k-1], [ $j+1, $k], [ $j+1, $k+1]) {
            my $aa = $_->[0];
            my $bb = $_->[1];
            if (getValue($aa,$bb) == 1){
            $count++;
            }

        }
        if ($inp_arr[$j][$k] == 0){
                if ($count == 2 || $count == 3) {
                    $print_arr[$j][$k] = 1;
                } else {
                    $print_arr[$j][$k] = 0;
		}
        }
        if ($inp_arr[$j][$k] == 1) {
                if ($count == 3) {
                    $print_arr[$j][$k] = 1;
                } else {
                    $print_arr[$j][$k] = 0;
        }
        }

	}
    }

    @inp_arr = map { [@$_] } @print_arr;
}
}
main()