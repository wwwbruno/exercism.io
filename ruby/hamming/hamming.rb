class Hamming
  def self.compute(dna1, dna2)
    raise ArgumentError if dna1.length != dna2.length

    dna1.chars.zip(dna2.chars).count { |x, y| x != y }
  end
end

module BookKeeping
  VERSION = 3
end
