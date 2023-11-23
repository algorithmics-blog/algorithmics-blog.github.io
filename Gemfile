source "https://rubygems.org"

gem "jekyll"
gem "webrick"
gem "github-pages", "~> 228", group: :jekyll_plugins

# If you have any plugins, put them here!
group :jekyll_plugins do
  gem "jekyll-feed", "~> 0.12"
  gem 'jekyll-paginate'
  gem "jekyll-tabs"
end

# Lock `http_parser.rb` gem to `v0.6.x` on JRuby builds since newer versions of the gem
# do not have a Java counterpart.
gem "http_parser.rb", "~> 0.6.0", :platforms => [:jruby]
